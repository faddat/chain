package types

import (
	"sort"
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidateBasic validates if a pool has a valid configuration
// meaning that it can be stored
func (pool *Pool) ValidateBasic(params Params) error {
	if pool.PoolId == UnknownID {
		return ErrInvalidID
	}
	if pool.State == PoolState_Unspecified {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "no state specified")
	}
	if err := sdk.ValidateDenom(pool.Denom); err != nil {
		return errorsmod.Wrapf(ErrInvalidPoolParams, err.Error())
	}
	if err := sdk.ValidateDenom(pool.NativeDenom); err != nil {
		return errorsmod.Wrapf(ErrInvalidPoolParams, err.Error())
	}
	if strings.TrimSpace(pool.ChainId) == "" {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "empty chain ID")
	}
	if strings.TrimSpace(pool.Bech32PrefixAccAddr) == "" {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "empty bech32 prefix account address")
	}
	if strings.TrimSpace(pool.Bech32PrefixValAddr) == "" {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "empty bech32 prefix validator address")
	}
	if len(pool.Validators) == 0 {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "empty validators set")
	} else {
		for _, val := range pool.Validators {
			bz, err := sdk.GetFromBech32(val.OperatorAddress, pool.Bech32PrefixValAddr)
			if err != nil {
				return errorsmod.Wrapf(ErrInvalidPoolParams, "invalid validator address %s: %v", val.OperatorAddress, err)
			}
			err = sdk.VerifyAddressFormat(bz)
			if err != nil {
				return errorsmod.Wrapf(ErrInvalidPoolParams, "invalid validator address %s: %v", val.OperatorAddress, err)
			}
		}
	}
	if pool.MinDepositAmount.IsNil() || pool.MinDepositAmount.LT(params.MinDepositAmount) {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "min deposit denom must be gte %d", params.MinDepositAmount.Int64())
	}
	if pool.AvailablePrizePool.IsNil() || pool.AvailablePrizePool.Denom != pool.Denom {
		return errorsmod.Wrapf(ErrInvalidPoolParams, "clawback prize pool must be initialized")
	}
	if err := pool.DrawSchedule.ValidateBasic(params); err != nil {
		return errorsmod.Wrapf(ErrInvalidPoolParams, err.Error())
	}
	if err := pool.PrizeStrategy.Validate(params); err != nil {
		return errorsmod.Wrapf(ErrInvalidPoolParams, err.Error())
	}
	return nil
}

func (p *Pool) IsLocalZone(ctx sdk.Context) bool {
	return p.ChainId == ctx.ChainID()
}

func (p *Pool) ShouldDraw(ctx sdk.Context) bool {
	if p.State == PoolState_Ready && (p.LastDrawState == DrawState_Unspecified || p.LastDrawState == DrawState_Success) {
		return p.DrawSchedule.ShouldDraw(ctx, p.LastDrawCreatedAt)
	}
	return false
}

// ActiveValidators returns currently enabled validators
func (p *Pool) ActiveValidators() (vals []*PoolValidator) {
	for _, v := range p.Validators {
		if v.IsEnabled {
			vals = append(vals, v)
		}
	}
	return
}

// BondedValidators returns active and inactive validators with a bonded amount > 0 for
func (p *Pool) BondedValidators() (activeVals, inactiveVals []*PoolValidator) {
	for _, v := range p.Validators {
		if v.IsBonded() {
			if v.IsEnabled {
				activeVals = append(activeVals, v)
			} else {
				inactiveVals = append(inactiveVals, v)
			}
		}
	}
	return
}

// ComputeSplitDelegations computes the delegation split to enforce based on the active validators in the set
// amount is divided evenly to all active validators
func (p *Pool) ComputeSplitDelegations(ctx sdk.Context, amount math.Int) (splits []*SplitDelegation) {
	activeValidators := p.ActiveValidators()
	if len(activeValidators) <= 0 {
		return nil
	}
	used := math.ZeroInt()
	valShare := amount.QuoRaw(int64(len(activeValidators)))

	for i, v := range activeValidators {
		// Compute the amount to use
		var relativeAmount math.Int
		if i == len(activeValidators)-1 {
			relativeAmount = amount.Sub(used)
		} else {
			relativeAmount = valShare
		}

		// Append to the destination structure
		splits = append(splits, &SplitDelegation{
			ValidatorAddress: v.OperatorAddress,
			Amount:           relativeAmount,
		})
		used = used.Add(relativeAmount)
	}

	if !used.Equal(amount) {
		// Returns nil in case we did something sketchy when computing amount
		return nil
	}

	return
}

// ComputeSplitUndelegations compute the undelegation split to enforce based on the bonded validators in the set
// disabled validators are prioritized and remaining amount is divided evenly between all validators
func (p *Pool) ComputeSplitUndelegations(ctx sdk.Context, amount math.Int) (splits []*SplitDelegation) {
	bondedActiveVals, bondedInactiveVals := p.BondedValidators()
	if len(bondedActiveVals) <= 0 && len(bondedInactiveVals) <= 0 {
		return nil
	}
	used := math.ZeroInt()

	// Sort vals by bonded amount descending to ensure we can fulfill the request
	sort.Slice(bondedInactiveVals, func(i, j int) bool {
		return bondedInactiveVals[i].BondedAmount.GT(bondedInactiveVals[j].BondedAmount)
	})
	for _, v := range bondedInactiveVals {
		// Undelegate as much as we can
		relativeAmount := math.MinInt(v.BondedAmount, amount.Sub(used))
		if relativeAmount.LTE(sdk.ZeroInt()) {
			continue
		}

		// Append to the destination structure
		splits = append(splits, &SplitDelegation{
			ValidatorAddress: v.OperatorAddress,
			Amount:           relativeAmount,
		})
		used = used.Add(relativeAmount)

		if used.Equal(amount) {
			break
		}
	}

	if !used.Equal(amount) {
		// Sort vals by bonded amount ascending to ensure we can fulfill the request
		sort.Slice(bondedActiveVals, func(i, j int) bool {
			return bondedActiveVals[j].BondedAmount.GT(bondedActiveVals[i].BondedAmount)
		})

		// Undelegate from active validator set
		for i, v := range bondedActiveVals {
			// Compute the amount to use each round to ensure fairness
			valShare := amount.Sub(used).QuoRaw(int64(len(bondedActiveVals) - i))
			relativeAmount := math.MinInt(v.BondedAmount, valShare)
			if i == len(bondedActiveVals)-1 && v.BondedAmount.GTE(amount.Sub(used)) {
				// Take remaining amount
				relativeAmount = amount.Sub(used)
			}
			if relativeAmount.LTE(sdk.ZeroInt()) {
				continue
			}

			// Append to the destination structure
			splits = append(splits, &SplitDelegation{
				ValidatorAddress: v.OperatorAddress,
				Amount:           relativeAmount,
			})
			used = used.Add(relativeAmount)

			if used.Equal(amount) {
				break
			}
		}
	}

	if !used.Equal(amount) {
		// Returns nil in case we cannot consume the whole amount
		return nil
	}

	return
}

func (p *Pool) ApplySplitDelegate(ctx sdk.Context, splits []*SplitDelegation) {
	for _, split := range splits {
		p.Validators[split.ValidatorAddress].BondedAmount = p.Validators[split.ValidatorAddress].BondedAmount.Add(split.Amount)
	}
}

func (p *Pool) ApplySplitUndelegate(ctx sdk.Context, splits []*SplitDelegation) {
	for _, split := range splits {
		p.Validators[split.ValidatorAddress].BondedAmount = p.Validators[split.ValidatorAddress].BondedAmount.Sub(split.Amount)
		if p.Validators[split.ValidatorAddress].BondedAmount.LT(sdk.ZeroInt()) {
			panic(ErrPoolInvalidSplit)
		}
	}
}