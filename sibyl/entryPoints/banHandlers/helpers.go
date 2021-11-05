package banHandlers

import sv "github.com/AnimeKaizoku/PsychoPass/sibyl/core/sibylValues"

func areAllSame(u *sv.User, banReason, banMsg, srcUrl string) bool {
	return u.Reason == banReason &&
		u.Message == banMsg &&
		u.BanSourceUrl == srcUrl && len(u.Message) > 0
}
