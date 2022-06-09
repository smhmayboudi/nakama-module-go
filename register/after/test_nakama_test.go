package after

import (
	"context"
	"os"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
)

type TestNakamaModule struct{}

// AuthenticateApple implements runtime.NakamaModule
func (nk *TestNakamaModule) AuthenticateApple(ctx context.Context, token, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateCustom implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateCustom(ctx context.Context, id, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateDevice implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateDevice(ctx context.Context, id, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateEmail implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateEmail(ctx context.Context, email, password, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateFacebook implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateFacebook(ctx context.Context, token string, importFriends bool, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateFacebookInstantGame implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateFacebookInstantGame(ctx context.Context, signedPlayerInfo string, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateGameCenter implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateGameCenter(ctx context.Context, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateGoogle implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateGoogle(ctx context.Context, token, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateSteam implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateSteam(ctx context.Context, token, username string, create bool) (string, string, bool, error) {
	return "", "", false, nil
}

// AuthenticateTokenGenerate implements tuntime.NakamaModule
func (nk *TestNakamaModule) AuthenticateTokenGenerate(userID, username string, exp int64, vars map[string]string) (string, int64, error) {
	return "", 0, nil
}

// AccountGetId implements tuntime.NakamaModule
func (nk *TestNakamaModule) AccountGetId(ctx context.Context, userID string) (*api.Account, error) {
	return nil, nil
}

// AccountsGetId implements tuntime.NakamaModule
func (nk *TestNakamaModule) AccountsGetId(ctx context.Context, userIDs []string) ([]*api.Account, error) {
	return nil, nil
}

// AccountUpdateId implements tuntime.NakamaModule
func (nk *TestNakamaModule) AccountUpdateId(ctx context.Context, userID, username string, metadata map[string]interface{}, displayName, timezone, location, langTag, avatarUrl string) error {
	return nil
}

// AccountDeleteId implements tuntime.NakamaModule
func (nk *TestNakamaModule) AccountDeleteId(ctx context.Context, userID string, recorded bool) error {
	return nil
}

// AccountExportId implements tuntime.NakamaModule
func (nk *TestNakamaModule) AccountExportId(ctx context.Context, userID string) (string, error) {
	return "", nil
}

// UsersGetId implements tuntime.NakamaModule
func (nk *TestNakamaModule) UsersGetId(ctx context.Context, userIDs []string, facebookIDs []string) ([]*api.User, error) {
	return nil, nil
}

// UsersGetUsername implements tuntime.NakamaModule
func (nk *TestNakamaModule) UsersGetUsername(ctx context.Context, usernames []string) ([]*api.User, error) {
	return nil, nil
}

// UsersGetRandom implements tuntime.NakamaModule
func (nk *TestNakamaModule) UsersGetRandom(ctx context.Context, count int) ([]*api.User, error) {
	return nil, nil
}

// UsersBanId implements tuntime.NakamaModule
func (nk *TestNakamaModule) UsersBanId(ctx context.Context, userIDs []string) error {
	return nil
}

// UsersUnbanId implements tuntime.NakamaModule
func (nk *TestNakamaModule) UsersUnbanId(ctx context.Context, userIDs []string) error {
	return nil
}

// LinkApple implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkApple(ctx context.Context, userID, token string) error {
	return nil
}

// LinkCustom implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkCustom(ctx context.Context, userID, customID string) error {
	return nil
}

// LinkDevice implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkDevice(ctx context.Context, userID, deviceID string) error {
	return nil
}

// LinkEmail implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkEmail(ctx context.Context, userID, email, password string) error {
	return nil
}

// LinkFacebook implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkFacebook(ctx context.Context, userID, username, token string, importFriends bool) error {
	return nil
}

// LinkFacebookInstantGame implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error {
	return nil
}

// LinkGameCenter implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error {
	return nil
}

// LinkGoogle implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkGoogle(ctx context.Context, userID, token string) error {
	return nil
}

// LinkSteam implements tuntime.NakamaModule
func (nk *TestNakamaModule) LinkSteam(ctx context.Context, userID, username, token string, importFriends bool) error {
	return nil
}

// ReadFile implements tuntime.NakamaModule
func (nk *TestNakamaModule) ReadFile(path string) (*os.File, error) {
	return nil, nil
}

// UnlinkApple implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkApple(ctx context.Context, userID, token string) error {
	return nil
}

// UnlinkCustom implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkCustom(ctx context.Context, userID, customID string) error {
	return nil
}

// UnlinkDevice implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkDevice(ctx context.Context, userID, deviceID string) error {
	return nil
}

// UnlinkEmail implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkEmail(ctx context.Context, userID, email string) error {
	return nil
}

// UnlinkFacebook implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkFacebook(ctx context.Context, userID, token string) error {
	return nil
}

// UnlinkFacebookInstantGame implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkFacebookInstantGame(ctx context.Context, userID, signedPlayerInfo string) error {
	return nil
}

// UnlinkGameCenter implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkGameCenter(ctx context.Context, userID, playerID, bundleID string, timestamp int64, salt, signature, publicKeyUrl string) error {
	return nil
}

// UnlinkGoogle implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkGoogle(ctx context.Context, userID, token string) error {
	return nil
}

// UnlinkSteam implements tuntime.NakamaModule
func (nk *TestNakamaModule) UnlinkSteam(ctx context.Context, userID, token string) error {
	return nil
}

// StreamUserList implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamUserList(mode uint8, subject, subcontext, label string, includeHidden, includeNotHidden bool) ([]runtime.Presence, error) {
	return nil, nil
}

// StreamUserGet implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamUserGet(mode uint8, subject, subcontext, label, userID, sessionID string) (runtime.PresenceMeta, error) {
	return nil, nil
}

// StreamUserJoin implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamUserJoin(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) (bool, error) {
	return false, nil
}

// StreamUserUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamUserUpdate(mode uint8, subject, subcontext, label, userID, sessionID string, hidden, persistence bool, status string) error {
	return nil
}

// StreamUserLeave implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamUserLeave(mode uint8, subject, subcontext, label, userID, sessionID string) error {
	return nil
}

// StreamUserKick implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamUserKick(mode uint8, subject, subcontext, label string, presence runtime.Presence) error {
	return nil
}

// StreamCount implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamCount(mode uint8, subject, subcontext, label string) (int, error) {
	return 0, nil
}

// StreamClose implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamClose(mode uint8, subject, subcontext, label string) error {
	return nil
}

// StreamSend implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamSend(mode uint8, subject, subcontext, label, data string, presences []runtime.Presence, reliable bool) error {
	return nil
}

// StreamSendRaw implements tuntime.NakamaModule
func (nk *TestNakamaModule) StreamSendRaw(mode uint8, subject, subcontext, label string, msg *rtapi.Envelope, presences []runtime.Presence, reliable bool) error {
	return nil
}

// SessionDisconnect implements tuntime.NakamaModule
func (nk *TestNakamaModule) SessionDisconnect(ctx context.Context, sessionID string, reason ...runtime.PresenceReason) error {
	return nil
}

// SessionLogout implements tuntime.NakamaModule
func (nk *TestNakamaModule) SessionLogout(userID, token, refreshToken string) error {
	return nil
}

// MatchCreate implements tuntime.NakamaModule
func (nk *TestNakamaModule) MatchCreate(ctx context.Context, module string, params map[string]interface{}) (string, error) {
	return "", nil
}

// MatchGet implements tuntime.NakamaModule
func (nk *TestNakamaModule) MatchGet(ctx context.Context, id string) (*api.Match, error) {
	return nil, nil
}

// MatchList implements tuntime.NakamaModule
func (nk *TestNakamaModule) MatchList(ctx context.Context, limit int, authoritative bool, label string, minSize, maxSize *int, query string) ([]*api.Match, error) {
	return nil, nil
}

// MatchSignal implements tuntime.NakamaModule
func (nk *TestNakamaModule) MatchSignal(ctx context.Context, id string, data string) (string, error) {
	return "", nil
}

// NotificationSend implements tuntime.NakamaModule
func (nk *TestNakamaModule) NotificationSend(ctx context.Context, userID, subject string, content map[string]interface{}, code int, sender string, persistent bool) error {
	return nil
}

// NotificationsSend implements tuntime.NakamaModule
func (nk *TestNakamaModule) NotificationsSend(ctx context.Context, notifications []*runtime.NotificationSend) error {
	return nil
}

// NotificationSendAll implements tuntime.NakamaModule
func (nk *TestNakamaModule) NotificationSendAll(ctx context.Context, subject string, content map[string]interface{}, code int, persistent bool) error {
	return nil
}

// WalletUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) WalletUpdate(ctx context.Context, userID string, changeset map[string]int64, metadata map[string]interface{}, updateLedger bool) (updated map[string]int64, previous map[string]int64, err error) {
	return nil, nil, nil
}

// WalletsUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) WalletsUpdate(ctx context.Context, updates []*runtime.WalletUpdate, updateLedger bool) ([]*runtime.WalletUpdateResult, error) {
	return nil, nil
}

// WalletLedgerUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) WalletLedgerUpdate(ctx context.Context, itemID string, metadata map[string]interface{}) (runtime.WalletLedgerItem, error) {
	return nil, nil
}

// WalletLedgerList implements tuntime.NakamaModule
func (nk *TestNakamaModule) WalletLedgerList(ctx context.Context, userID string, limit int, cursor string) ([]runtime.WalletLedgerItem, string, error) {
	return nil, "", nil
}

// StorageList implements tuntime.NakamaModule
func (nk *TestNakamaModule) StorageList(ctx context.Context, userID, collection string, limit int, cursor string) ([]*api.StorageObject, string, error) {
	return nil, "", nil
}

// StorageRead implements tuntime.NakamaModule
func (nk *TestNakamaModule) StorageRead(ctx context.Context, reads []*runtime.StorageRead) ([]*api.StorageObject, error) {
	return nil, nil
}

// StorageWrite implements tuntime.NakamaModule
func (nk *TestNakamaModule) StorageWrite(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
	return nil, nil
}

// StorageDelete implements tuntime.NakamaModule
func (nk *TestNakamaModule) StorageDelete(ctx context.Context, deletes []*runtime.StorageDelete) error {
	return nil
}

// MultiUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) MultiUpdate(ctx context.Context, accountUpdates []*runtime.AccountUpdate, storageWrites []*runtime.StorageWrite, walletUpdates []*runtime.WalletUpdate, updateLedger bool) ([]*api.StorageObjectAck, []*runtime.WalletUpdateResult, error) {
	return nil, nil, nil
}

// LeaderboardCreate implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}) error {
	return nil
}

// LeaderboardDelete implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardDelete(ctx context.Context, id string) error {
	return nil
}

// LeaderboardList implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardList(categoryStart, categoryEnd, limit int, cursor string) (*api.LeaderboardList, error) {
	return nil, nil
}

// LeaderboardRecordsList implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardRecordsList(ctx context.Context, id string, ownerIDs []string, limit int, cursor string, expiry int64) (records []*api.LeaderboardRecord, ownerRecords []*api.LeaderboardRecord, nextCursor string, prevCursor string, err error) {
	return nil, nil, "", "", nil
}

// LeaderboardRecordWrite implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}, overrideOperator *int) (*api.LeaderboardRecord, error) {
	return nil, nil
}

// LeaderboardRecordDelete implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardRecordDelete(ctx context.Context, id, ownerID string) error {
	return nil
}

// LeaderboardsGetId implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardsGetId(ctx context.Context, ids []string) ([]*api.Leaderboard, error) {
	return nil, nil
}

// LeaderboardRecordsHaystack implements tuntime.NakamaModule
func (nk *TestNakamaModule) LeaderboardRecordsHaystack(ctx context.Context, id, ownerID string, limit int, expiry int64) ([]*api.LeaderboardRecord, error) {
	return nil, nil
}

// PurchaseValidateApple implements tuntime.NakamaModule
func (nk *TestNakamaModule) PurchaseValidateApple(ctx context.Context, userID, receipt string, persist bool, passwordOverride ...string) (*api.ValidatePurchaseResponse, error) {
	return nil, nil
}

// PurchaseValidateGoogle implements tuntime.NakamaModule
func (nk *TestNakamaModule) PurchaseValidateGoogle(ctx context.Context, userID, receipt string, persist bool, overrides ...struct {
	ClientEmail string
	PrivateKey  string
}) (*api.ValidatePurchaseResponse, error) {
	return nil, nil
}

// PurchaseValidateHuawei implements tuntime.NakamaModule
func (nk *TestNakamaModule) PurchaseValidateHuawei(ctx context.Context, userID, signature, inAppPurchaseData string, persist bool) (*api.ValidatePurchaseResponse, error) {
	return nil, nil
}

// PurchasesList implements tuntime.NakamaModule
func (nk *TestNakamaModule) PurchasesList(ctx context.Context, userID string, limit int, cursor string) (*api.PurchaseList, error) {
	return nil, nil
}

// PurchaseGetByTransactionId implements tuntime.NakamaModule
func (nk *TestNakamaModule) PurchaseGetByTransactionId(ctx context.Context, transactionID string) (string, *api.ValidatedPurchase, error) {
	return "", nil, nil
}

// TournamentCreate implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentCreate(ctx context.Context, id string, authoritative bool, sortOrder, operator, resetSchedule string, metadata map[string]interface{}, title, description string, category, startTime, endTime, duration, maxSize, maxNumScore int, joinRequired bool) error {
	return nil
}

// TournamentDelete implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentDelete(ctx context.Context, id string) error {
	return nil
}

// TournamentAddAttempt implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentAddAttempt(ctx context.Context, id, ownerID string, count int) error {
	return nil
}

// TournamentJoin implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentJoin(ctx context.Context, id, ownerID, username string) error {
	return nil
}

// TournamentsGetId implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentsGetId(ctx context.Context, tournamentIDs []string) ([]*api.Tournament, error) {
	return nil, nil
}

// TournamentList implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentList(ctx context.Context, categoryStart, categoryEnd, startTime, endTime, limit int, cursor string) (*api.TournamentList, error) {
	return nil, nil
}

// TournamentRecordsList implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentRecordsList(ctx context.Context, tournamentId string, ownerIDs []string, limit int, cursor string, overrideExpiry int64) (records []*api.LeaderboardRecord, ownerRecords []*api.LeaderboardRecord, prevCursor string, nextCursor string, err error) {
	return nil, nil, "", "", nil
}

// TournamentRecordWrite implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentRecordWrite(ctx context.Context, id, ownerID, username string, score, subscore int64, metadata map[string]interface{}, operatorOverride *int) (*api.LeaderboardRecord, error) {
	return nil, nil
}

// TournamentRecordsHaystack implements tuntime.NakamaModule
func (nk *TestNakamaModule) TournamentRecordsHaystack(ctx context.Context, id, ownerID string, limit int, expiry int64) ([]*api.LeaderboardRecord, error) {
	return nil, nil
}

// GroupsGetId implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupsGetId(ctx context.Context, groupIDs []string) ([]*api.Group, error) {
	return nil, nil
}

// GroupCreate implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupCreate(ctx context.Context, userID, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) (*api.Group, error) {
	return nil, nil
}

// GroupUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUpdate(ctx context.Context, id, name, creatorID, langTag, description, avatarUrl string, open bool, metadata map[string]interface{}, maxCount int) error {
	return nil
}

// GroupDelete implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupDelete(ctx context.Context, id string) error {
	return nil
}

// GroupUserJoin implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUserJoin(ctx context.Context, groupID, userID, username string) error {
	return nil
}

// GroupUserLeave implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUserLeave(ctx context.Context, groupID, userID, username string) error {
	return nil
}

// GroupUsersAdd implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUsersAdd(ctx context.Context, callerID, groupID string, userIDs []string) error {
	return nil
}

// GroupUsersBan implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUsersBan(ctx context.Context, callerID, groupID string, userIDs []string) error {
	return nil
}

// GroupUsersKick implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUsersKick(ctx context.Context, callerID, groupID string, userIDs []string) error {
	return nil
}

// GroupUsersPromote implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUsersPromote(ctx context.Context, callerID, groupID string, userIDs []string) error {
	return nil
}

// GroupUsersDemote implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUsersDemote(ctx context.Context, callerID, groupID string, userIDs []string) error {
	return nil
}

// GroupUsersList implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupUsersList(ctx context.Context, id string, limit int, state *int, cursor string) ([]*api.GroupUserList_GroupUser, string, error) {
	return nil, "", nil
}

// GroupsList implements tuntime.NakamaModule
func (nk *TestNakamaModule) GroupsList(ctx context.Context, name, langTag string, members *int, open *bool, limit int, cursor string) ([]*api.Group, string, error) {
	return nil, "", nil
}

// UserGroupsList implements tuntime.NakamaModule
func (nk *TestNakamaModule) UserGroupsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.UserGroupList_UserGroup, string, error) {
	return nil, "", nil
}

// FriendsList implements tuntime.NakamaModule
func (nk *TestNakamaModule) FriendsList(ctx context.Context, userID string, limit int, state *int, cursor string) ([]*api.Friend, string, error) {
	return nil, "", nil
}

// FriendsAdd implements tuntime.NakamaModule
func (nk *TestNakamaModule) FriendsAdd(ctx context.Context, userID string, username string, ids []string, usernames []string) error {
	return nil
}

// FriendsDelete implements tuntime.NakamaModule
func (nk *TestNakamaModule) FriendsDelete(ctx context.Context, userID string, username string, ids []string, usernames []string) error {
	return nil
}

// FriendsBlock implements tuntime.NakamaModule
func (nk *TestNakamaModule) FriendsBlock(ctx context.Context, userID string, username string, ids []string, usernames []string) error {
	return nil
}

// Event implements tuntime.NakamaModule
func (nk *TestNakamaModule) Event(ctx context.Context, evt *api.Event) error {
	return nil
}

// MetricsCounterAdd implements tuntime.NakamaModule
func (nk *TestNakamaModule) MetricsCounterAdd(name string, tags map[string]string, delta int64) {
}

// MetricsGaugeSet implements tuntime.NakamaModule
func (nk *TestNakamaModule) MetricsGaugeSet(name string, tags map[string]string, value float64) {
}

// MetricsTimerRecord implements tuntime.NakamaModule
func (nk *TestNakamaModule) MetricsTimerRecord(name string, tags map[string]string, value time.Duration) {
}

// ChannelIdBuild implements tuntime.NakamaModule
func (nk *TestNakamaModule) ChannelIdBuild(ctx context.Context, sender string, target string, chanType runtime.ChannelType) (string, error) {
	return "", nil
}

// ChannelMessageSend implements tuntime.NakamaModule
func (nk *TestNakamaModule) ChannelMessageSend(ctx context.Context, channelID string, content map[string]interface{}, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error) {
	return nil, nil
}

// ChannelMessageUpdate implements tuntime.NakamaModule
func (nk *TestNakamaModule) ChannelMessageUpdate(ctx context.Context, channelID, messageID string, content map[string]interface{}, senderId, senderUsername string, persist bool) (*rtapi.ChannelMessageAck, error) {
	return nil, nil
}

// ChannelMessagesList implements tuntime.NakamaModule
func (nk *TestNakamaModule) ChannelMessagesList(ctx context.Context, channelId string, limit int, forward bool, cursor string) (messages []*api.ChannelMessage, nextCursor string, prevCursor string, err error) {
	return nil, "", "", nil
}

var _ runtime.NakamaModule = (*TestNakamaModule)(nil)
