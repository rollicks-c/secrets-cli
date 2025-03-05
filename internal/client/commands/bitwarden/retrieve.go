package bitwarden

import (
	"fmt"
	"github.com/rollicks-c/secrets-cli/internal/logic/ui"
	"github.com/rollicks-c/term"
)

func getSecret(itemID, key string, toClipboard bool) error {

	// load from bitwarden
	client, err := connect()
	if err != nil {
		return err
	}
	item, err := client.GetItem(itemID)
	if err != nil {
		return err
	}

	// extract
	key, value, ok := lookupKey(item, key)
	if !ok {
		return fmt.Errorf("key [%s] not found in item [%s]", key, item.Name)
	}

	// provide
	if toClipboard {
		if err := ui.PutInClipboard(value); err != nil {
			return err
		}
		notification := fmt.Sprintf("BW secret [%s] of [%s] ready in clipboard", key, item.Name)
		if err := ui.NotifyUser(notification); err != nil {
			return err
		}
		term.Successf("Secret [%s] of [%s] ready in clipboard\n", key, item.Name)
	} else {
		fmt.Print(value)
	}

	return nil
}

func getTOTP(itemID string, toClipboard bool) error {

	// load from bitwarden
	client, err := connect()
	if err != nil {
		return err
	}
	item, err := client.GetTOTP(itemID)
	if err != nil {
		return err
	}

	// provide
	if toClipboard {
		if err := ui.PutInClipboard(item.Data); err != nil {
			return err
		}
		notification := fmt.Sprintf("BW TOTP [%s] ready in clipboard", itemID)
		if err := ui.NotifyUser(notification); err != nil {
			return err
		}
		term.Successf("TOTP [%s] ready in clipboard\n", itemID)
	} else {
		fmt.Print(item.Data)
	}

	return nil
}
