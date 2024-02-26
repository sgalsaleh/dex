// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/oauth2client"
)

// OAuth2Client is the model entity for the OAuth2Client schema.
type OAuth2Client struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"secret,omitempty"`
	// RedirectUris holds the value of the "redirect_uris" field.
	RedirectUris []string `json:"redirect_uris,omitempty"`
	// TrustedPeers holds the value of the "trusted_peers" field.
	TrustedPeers []string `json:"trusted_peers,omitempty"`
	// Public holds the value of the "public" field.
	Public bool `json:"public,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL      string `json:"logo_url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuth2Client) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauth2client.FieldRedirectUris, oauth2client.FieldTrustedPeers:
			values[i] = new([]byte)
		case oauth2client.FieldPublic:
			values[i] = new(sql.NullBool)
		case oauth2client.FieldID, oauth2client.FieldSecret, oauth2client.FieldName, oauth2client.FieldLogoURL:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuth2Client fields.
func (o *OAuth2Client) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauth2client.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				o.ID = value.String
			}
		case oauth2client.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				o.Secret = value.String
			}
		case oauth2client.FieldRedirectUris:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uris", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.RedirectUris); err != nil {
					return fmt.Errorf("unmarshal field redirect_uris: %w", err)
				}
			}
		case oauth2client.FieldTrustedPeers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field trusted_peers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.TrustedPeers); err != nil {
					return fmt.Errorf("unmarshal field trusted_peers: %w", err)
				}
			}
		case oauth2client.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				o.Public = value.Bool
			}
		case oauth2client.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		case oauth2client.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo_url", values[i])
			} else if value.Valid {
				o.LogoURL = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OAuth2Client.
// This includes values selected through modifiers, order, etc.
func (o *OAuth2Client) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// Update returns a builder for updating this OAuth2Client.
// Note that you need to call OAuth2Client.Unwrap() before calling this method if this OAuth2Client
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *OAuth2Client) Update() *OAuth2ClientUpdateOne {
	return NewOAuth2ClientClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the OAuth2Client entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *OAuth2Client) Unwrap() *OAuth2Client {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("db: OAuth2Client is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *OAuth2Client) String() string {
	var builder strings.Builder
	builder.WriteString("OAuth2Client(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("secret=")
	builder.WriteString(o.Secret)
	builder.WriteString(", ")
	builder.WriteString("redirect_uris=")
	builder.WriteString(fmt.Sprintf("%v", o.RedirectUris))
	builder.WriteString(", ")
	builder.WriteString("trusted_peers=")
	builder.WriteString(fmt.Sprintf("%v", o.TrustedPeers))
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", o.Public))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(o.LogoURL)
	builder.WriteByte(')')
	return builder.String()
}

// OAuth2Clients is a parsable slice of OAuth2Client.
type OAuth2Clients []*OAuth2Client
