package testutil

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/opzlabs/cosmos-sdk/client"
	"github.com/opzlabs/cosmos-sdk/testutil"
	clitestutil "github.com/opzlabs/cosmos-sdk/testutil/cli"
	sdk "github.com/opzlabs/cosmos-sdk/types"
	bankcli "github.com/opzlabs/cosmos-sdk/x/bank/client/cli"
)

func MsgSendExec(clientCtx client.Context, from, to, amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{from.String(), to.String(), amount.String()}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, bankcli.NewSendTxCmd(), args)
}

func MsgMultiSendExec(clientCtx client.Context, from sdk.AccAddress, to []sdk.AccAddress, amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{from.String()}
	for _, addr := range to {
		args = append(args, addr.String())
	}

	args = append(args, amount.String())
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, bankcli.NewMultiSendTxCmd(), args)
}

func QueryBalancesExec(clientCtx client.Context, address fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{address.String(), fmt.Sprintf("--%s=json", cli.OutputFlag)}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, bankcli.GetBalancesCmd(), args)
}
