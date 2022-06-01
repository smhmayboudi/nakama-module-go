package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	r "github.com/smhmayboudi/nakama-modules-go/register"
	ra "github.com/smhmayboudi/nakama-modules-go/register/after"
	rb "github.com/smhmayboudi/nakama-modules-go/register/before"
	u "github.com/smhmayboudi/nakama-modules-go/util"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	u.NewConfig(ctx, logger)

	u.NewOpenTelemetry(ctx, logger)
	// shutdown := u.NewOpenTelemetry(ctx, logger)
	// defer shutdown()

	ctx = u.Extract(ctx, b3.B3SingleHeader)
	nakamaContext := u.NewContext(ctx, logger)
	fields := map[string]interface{}{"name": "InitModule", "ctx": nakamaContext}
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(fields).Debug("")
	_, span := otel.Tracer(u.AppConfig.InstrumentationName).Start(
		ctx,
		"InitModule",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := initializer.RegisterAfterAddFriends(ra.RegisterAfterAddFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAddGroupUsers(ra.RegisterAfterAddGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateApple(ra.RegisterAfterAuthenticateApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateCustom(ra.RegisterAfterAuthenticateCustom); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateDevice(ra.RegisterAfterAuthenticateDevice); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateEmail(ra.RegisterAfterAuthenticateEmail); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateFacebook(ra.RegisterAfterAuthenticateFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateFacebookInstantGame(ra.RegisterAfterAuthenticateFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateGameCenter(ra.RegisterAfterAuthenticateGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateGoogle(ra.RegisterAfterAuthenticateGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateSteam(ra.RegisterAfterAuthenticateSteam); err != nil {
		return err
	}
	if err := initializer.RegisterAfterBanGroupUsers(ra.RegisterAfterBanGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterBlockFriends(ra.RegisterAfterBlockFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterCreateGroup(ra.RegisterAfterCreateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteFriends(ra.RegisterAfterDeleteFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteGroup(ra.RegisterAfterDeleteGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteLeaderboardRecord(ra.RegisterAfterDeleteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteNotifications(ra.RegisterAfterDeleteNotifications); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteStorageObjects(ra.RegisterAfterDeleteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDemoteGroupUsers(ra.RegisterAfterDemoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterGetAccount(ra.RegisterAfterGetAccount); err != nil {
		return err
	}
	if err := initializer.RegisterAfterGetUsers(ra.RegisterAfterGetUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterImportFacebookFriends(ra.RegisterAfterImportFacebookFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterImportSteamFriends(ra.RegisterAfterImportSteamFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterJoinGroup(ra.RegisterAfterJoinGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterJoinTournament(ra.RegisterAfterJoinTournament); err != nil {
		return err
	}
	if err := initializer.RegisterAfterKickGroupUsers(ra.RegisterAfterKickGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLeaveGroup(ra.RegisterAfterLeaveGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkApple(ra.RegisterAfterLinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkCustom(ra.RegisterAfterLinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkDevice(ra.RegisterAfterLinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkEmail(ra.RegisterAfterLinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkFacebook(ra.RegisterAfterLinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkFacebookInstantGame(ra.RegisterAfterLinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkGameCenter(ra.RegisterAfterLinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkGoogle(ra.RegisterAfterLinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkSteam(ra.RegisterAfterLinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListChannelMessages(ra.RegisterAfterListChannelMessages); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListFriends(ra.RegisterAfterListFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListGroupUsers(ra.RegisterAfterListGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListGroups(ra.RegisterAfterListGroups); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListLeaderboardRecords(ra.RegisterAfterListLeaderboardRecords); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListLeaderboardRecordsAroundOwner(ra.RegisterAfterListLeaderboardRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListMatches(ra.RegisterAfterListMatches); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListNotifications(ra.RegisterAfterListNotifications); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListStorageObjects(ra.RegisterAfterListStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListTournamentRecords(ra.RegisterAfterListTournamentRecords); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListTournamentRecordsAroundOwner(ra.RegisterAfterListTournamentRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListTournaments(ra.RegisterAfterListTournaments); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListUserGroups(ra.RegisterAfterListUserGroups); err != nil {
		return err
	}
	if err := initializer.RegisterAfterPromoteGroupUsers(ra.RegisterAfterPromoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterReadStorageObjects(ra.RegisterAfterReadStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterSessionLogout(ra.RegisterAfterSessionLogout); err != nil {
		return err
	}
	if err := initializer.RegisterAfterSessionRefresh(ra.RegisterAfterSessionRefresh); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkApple(ra.RegisterAfterUnlinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkCustom(ra.RegisterAfterUnlinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkDevice(ra.RegisterAfterUnlinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkEmail(ra.RegisterAfterUnlinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkFacebook(ra.RegisterAfterUnlinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkFacebookInstantGame(ra.RegisterAfterUnlinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkGameCenter(ra.RegisterAfterUnlinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkGoogle(ra.RegisterAfterUnlinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkSteam(ra.RegisterAfterUnlinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUpdateAccount(ra.RegisterAfterUpdateAccount); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUpdateGroup(ra.RegisterAfterUpdateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterValidatePurchaseApple(ra.RegisterAfterValidatePurchaseApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterValidatePurchaseGoogle(ra.RegisterAfterValidatePurchaseGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterValidatePurchaseHuawei(ra.RegisterAfterValidatePurchaseHuawei); err != nil {
		return err
	}
	if err := initializer.RegisterAfterWriteLeaderboardRecord(ra.RegisterAfterWriteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterAfterWriteStorageObjects(ra.RegisterAfterWriteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterWriteTournamentRecord(ra.RegisterAfterWriteTournamentRecord); err != nil {
		return err
	}

	if err := initializer.RegisterBeforeAddFriends(rb.RegisterBeforeAddFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAddGroupUsers(rb.RegisterBeforeAddGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateApple(rb.RegisterBeforeAuthenticateApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateCustom(rb.RegisterBeforeAuthenticateCustom); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateDevice(rb.RegisterBeforeAuthenticateDevice); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateEmail(rb.RegisterBeforeAuthenticateEmail); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateFacebook(rb.RegisterBeforeAuthenticateFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateFacebookInstantGame(rb.RegisterBeforeAuthenticateFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateGameCenter(rb.RegisterBeforeAuthenticateGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateGoogle(rb.RegisterBeforeAuthenticateGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateSteam(rb.RegisterBeforeAuthenticateSteam); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeBanGroupUsers(rb.RegisterBeforeBanGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeBlockFriends(rb.RegisterBeforeBlockFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeCreateGroup(rb.RegisterBeforeCreateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteFriends(rb.RegisterBeforeDeleteFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteGroup(rb.RegisterBeforeDeleteGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteLeaderboardRecord(rb.RegisterBeforeDeleteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteNotifications(rb.RegisterBeforeDeleteNotifications); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteStorageObjects(rb.RegisterBeforeDeleteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDemoteGroupUsers(rb.RegisterBeforeDemoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeGetAccount(rb.RegisterBeforeGetAccount); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeGetUsers(rb.RegisterBeforeGetUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeImportFacebookFriends(rb.RegisterBeforeImportFacebookFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeImportSteamFriends(rb.RegisterBeforeImportSteamFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeJoinGroup(rb.RegisterBeforeJoinGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeJoinTournament(rb.RegisterBeforeJoinTournament); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeKickGroupUsers(rb.RegisterBeforeKickGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLeaveGroup(rb.RegisterBeforeLeaveGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkApple(rb.RegisterBeforeLinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkCustom(rb.RegisterBeforeLinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkDevice(rb.RegisterBeforeLinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkEmail(rb.RegisterBeforeLinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkFacebook(rb.RegisterBeforeLinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkFacebookInstantGame(rb.RegisterBeforeLinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkGameCenter(rb.RegisterBeforeLinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkGoogle(rb.RegisterBeforeLinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkSteam(rb.RegisterBeforeLinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListChannelMessages(rb.RegisterBeforeListChannelMessages); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListFriends(rb.RegisterBeforeListFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListGroupUsers(rb.RegisterBeforeListGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListGroups(rb.RegisterBeforeListGroups); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListLeaderboardRecords(rb.RegisterBeforeListLeaderboardRecords); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListLeaderboardRecordsAroundOwner(rb.RegisterBeforeListLeaderboardRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListMatches(rb.RegisterBeforeListMatches); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListNotifications(rb.RegisterBeforeListNotifications); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListStorageObjects(rb.RegisterBeforeListStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListTournamentRecords(rb.RegisterBeforeListTournamentRecords); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListTournamentRecordsAroundOwner(rb.RegisterBeforeListTournamentRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListTournaments(rb.RegisterBeforeListTournaments); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListUserGroups(rb.RegisterBeforeListUserGroups); err != nil {
		return err
	}
	if err := initializer.RegisterBeforePromoteGroupUsers(rb.RegisterBeforePromoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeReadStorageObjects(rb.RegisterBeforeReadStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeSessionLogout(rb.RegisterBeforeSessionLogout); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeSessionRefresh(rb.RegisterBeforeSessionRefresh); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkApple(rb.RegisterBeforeUnlinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkCustom(rb.RegisterBeforeUnlinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkDevice(rb.RegisterBeforeUnlinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkEmail(rb.RegisterBeforeUnlinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkFacebook(rb.RegisterBeforeUnlinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkFacebookInstantGame(rb.RegisterBeforeUnlinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkGameCenter(rb.RegisterBeforeUnlinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkGoogle(rb.RegisterBeforeUnlinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkSteam(rb.RegisterBeforeUnlinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUpdateAccount(rb.RegisterBeforeUpdateAccount); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUpdateGroup(rb.RegisterBeforeUpdateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeValidatePurchaseApple(rb.RegisterBeforeValidatePurchaseApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeValidatePurchaseGoogle(rb.RegisterBeforeValidatePurchaseGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeValidatePurchaseHuawei(rb.RegisterBeforeValidatePurchaseHuawei); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeWriteLeaderboardRecord(rb.RegisterBeforeWriteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeWriteStorageObjects(rb.RegisterBeforeWriteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeWriteTournamentRecord(rb.RegisterBeforeWriteTournamentRecord); err != nil {
		return err
	}

	if err := initializer.RegisterAfterRt("ChannelJoin", ra.RegisterAfterChannelJoin); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelLeave", ra.RegisterAfterChannelLeave); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelMessageRemove", ra.RegisterAfterChannelMessageRemove); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelMessageSend", ra.RegisterAfterChannelMessageSend); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelMessageUpdate", ra.RegisterAfterChannelMessageUpdate); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchCreate", ra.RegisterAfterMatchCreate); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchDataSend", ra.RegisterAfterMatchDataSend); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchJoin", ra.RegisterAfterMatchJoin); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchLeave", ra.RegisterAfterMatchLeave); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchmakerAdd", ra.RegisterAfterMatchmakerAdd); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchmakerRemove", ra.RegisterAfterMatchmakerRemove); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("Ping", ra.RegisterAfterPing); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("Pong", ra.RegisterAfterPong); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("StatusFollow", ra.RegisterAfterStatusFollow); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("StatusUnfollow", ra.RegisterAfterStatusUnfollow); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("StatusUpdate", ra.RegisterAfterStatusUpdate); err != nil {
		return err
	}

	if err := initializer.RegisterBeforeRt("ChannelJoin", rb.RegisterBeforeChannelJoin); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelLeave", rb.RegisterBeforeChannelLeave); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelMessageRemove", rb.RegisterBeforeChannelMessageRemove); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelMessageSend", rb.RegisterBeforeChannelMessageSend); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelMessageUpdate", rb.RegisterBeforeChannelMessageUpdate); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchCreate", rb.RegisterBeforeMatchCreate); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchDataSend", rb.RegisterBeforeMatchDataSend); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchJoin", rb.RegisterBeforeMatchJoin); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchLeave", rb.RegisterBeforeMatchLeave); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchmakerAdd", rb.RegisterBeforeMatchmakerAdd); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchmakerRemove", rb.RegisterBeforeMatchmakerRemove); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("Ping", rb.RegisterBeforePing); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("Pong", rb.RegisterBeforePong); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("StatusFollow", rb.RegisterBeforeStatusFollow); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("StatusUnfollow", rb.RegisterBeforeStatusUnfollow); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("StatusUpdate", rb.RegisterBeforeStatusUpdate); err != nil {
		return err
	}
	if err := initializer.RegisterEvent(r.RegisterEvent); err != nil {
		return err
	}
	if err := initializer.RegisterEventSessionEnd(r.RegisterEventSessionEnd); err != nil {
		return err
	}
	if err := initializer.RegisterEventSessionStart(r.RegisterEventSessionStart); err != nil {
		return err
	}
	if err := initializer.RegisterLeaderboardReset(r.RegisterLeaderboardReset); err != nil {
		return err
	}
	if err := initializer.RegisterMatch("", r.RegisterMatch); err != nil {
		return err
	}
	if err := initializer.RegisterMatchmakerMatched(r.RegisterMatchmakerMatched); err != nil {
		return err
	}
	if err := initializer.RegisterTournamentEnd(r.RegisterTournamentEnd); err != nil {
		return err
	}
	if err := initializer.RegisterTournamentReset(r.RegisterTournamentReset); err != nil {
		return err
	}
	return nil
}
