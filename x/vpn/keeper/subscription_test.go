package keeper

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_SetSubscriptionsCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCount(ctx, 1)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCount(ctx, 2)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(2), count)

	keeper.SetSubscriptionsCount(ctx, 3)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(3), count)
}

func TestKeeper_GetSubscriptionsCount(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCount(ctx, 1)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCount(ctx, 2)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(2), count)

	keeper.SetSubscriptionsCount(ctx, 3)
	count = keeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(3), count)
}

func TestKeeper_SetSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)

	keeper.SetSubscription(ctx, types.TestSubscriptionEmpty)
	result, found := keeper.GetSubscription(ctx, types.TestSubscriptionEmpty.ID) //TODO empty subscription should not be set
	require.Equal(t, true, found)

	keeper.SetSubscription(ctx, types.TestSubscriptionValid)
	result, found = keeper.GetSubscription(ctx, types.TestNodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSubscriptionValid, result)
}

func TestKeeper_GetSubscription(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)

	keeper.SetSubscription(ctx, types.TestSubscriptionEmpty)
	result, found := keeper.GetSubscription(ctx, types.TestSubscriptionEmpty.ID) //TODO empty subscription should not be set
	require.Equal(t, true, found)

	keeper.SetSubscription(ctx, types.TestSubscriptionValid)
	result, found = keeper.GetSubscription(ctx, types.TestNodeValid.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSubscriptionValid, result)
}

func TestKeeper_SetSubscriptionsCountOfNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 1)
	count = keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 2)
	count = keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(2), count)

	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 1)
	count = keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSubscriptionsCountOfNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 1)
	count = keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 2)
	count = keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(2), count)

	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 1)
	count = keeper.GetSubscriptionsCountOfNode(ctx, types.TestIDZero)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetSubscriptionIDByNodeID(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDZero, 1)
	require.Equal(t, false, found)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDZero, 1, types.TestIDZero)
	id, found := keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDZero, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDZero, 2, types.TestIDPos)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDZero, 2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDPos, 1, types.TestIDZero)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDPos, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDPos, 2, types.TestIDPos)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDPos, 2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)
}

func TestKeeper_GetSubscriptionIDByNodeID(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDZero, 1)
	require.Equal(t, false, found)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDZero, 1, types.TestIDZero)
	id, found := keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDZero, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDZero, 2, types.TestIDPos)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDZero, 2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDPos, 1, types.TestIDZero)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDPos, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSubscriptionIDByNodeID(ctx, types.TestIDPos, 2, types.TestIDPos)
	id, found = keeper.GetSubscriptionIDByNodeID(ctx, types.TestIDPos, 2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)
}

func TestKeeper_SetSubscriptionsCountOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 2)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(2), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_GetSubscriptionsCountOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	count := keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 2)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(2), count)

	keeper.SetSubscriptionsCountOfAddress(ctx, types.TestAddress1, 1)
	count = keeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)
}

func TestKeeper_SetSubscriptionIDByAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, false, found)

	keeper.SetSubscriptionIDByAddress(ctx, types.TestAddressEmpty, 1, types.TestIDZero) // TODO empty address should not be set
	id, found := keeper.GetSubscriptionIDByAddress(ctx, types.TestAddressEmpty, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 1, types.TestIDPos)
	id, found = keeper.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

	keeper.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 1, types.TestIDZero)
	id, found = keeper.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)
}

func TestKeeper_GetSubscriptionIDByAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	_, found := keeper.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, false, found)

	keeper.SetSubscriptionIDByAddress(ctx, types.TestAddressEmpty, 1, types.TestIDZero) // TODO empty address should not be set
	id, found := keeper.GetSubscriptionIDByAddress(ctx, types.TestAddressEmpty, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)

	keeper.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 1, types.TestIDPos)
	id, found = keeper.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, id)

	keeper.SetSubscriptionIDByAddress(ctx, types.TestAddress1, 1, types.TestIDZero)
	id, found = keeper.GetSubscriptionIDByAddress(ctx, types.TestAddress1, 1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, id)
}

func TestKeeper_GetSubscriptionsOfNode(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	subscriptions := keeper.GetSubscriptionsOfNode(ctx, types.TestIDZero)
	require.Equal(t, types.TestSubscriptionsEmpty, subscriptions)

	keeper.SetSubscription(ctx, types.TestSubscriptionValid)
	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 1)

	subscriptions = keeper.GetSubscriptionsOfNode(ctx, types.TestIDZero)
	require.Equal(t, types.TestSubscriptionsValid, subscriptions)

	subscription := types.TestSubscriptionValid // TODO same subscription should not be store
	keeper.SetSubscription(ctx, subscription)
	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 2)

	subscriptions = keeper.GetSubscriptionsOfNode(ctx, types.TestIDZero)
	require.Equal(t, append(types.TestSubscriptionsValid, subscription), subscriptions)

}

func TestKeeper_GetSubscriptionsOfAddress(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	subscriptions := keeper.GetSubscriptionsOfNode(ctx, types.TestIDZero)
	require.Equal(t, types.TestSubscriptionsEmpty, subscriptions)

	keeper.SetSubscription(ctx, types.TestSubscriptionValid)
	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 1)

	subscriptions = keeper.GetSubscriptionsOfNode(ctx, types.TestIDZero)
	require.Equal(t, types.TestSubscriptionsValid, subscriptions)

	subscription := types.TestSubscriptionValid // TODO same subscription should not be store
	keeper.SetSubscription(ctx, subscription)
	keeper.SetSubscriptionsCountOfNode(ctx, types.TestIDZero, 2)

	subscriptions = keeper.GetSubscriptionsOfNode(ctx, types.TestIDZero)
	require.Equal(t, append(types.TestSubscriptionsValid, subscription), subscriptions)
}

func TestKeeper_GetAllSubscriptions(t *testing.T) {
	ctx, _, _, keeper, _, _ := TestCreateInput()

	subscriptions := keeper.GetAllSubscriptions(ctx)
	require.Equal(t, types.TestSubscriptionsNil, subscriptions)

	keeper.SetSubscription(ctx, types.TestSubscriptionValid)
	subscriptions = keeper.GetAllSubscriptions(ctx)
	require.Equal(t, types.TestSubscriptionsValid, subscriptions)

	subscription := types.TestSubscriptionValid
	subscription.ID = types.TestIDPos
	keeper.SetSubscription(ctx, subscription)
	subscriptions = keeper.GetAllSubscriptions(ctx)
	require.Equal(t, append(types.TestSubscriptionsValid, subscription), subscriptions)
}
