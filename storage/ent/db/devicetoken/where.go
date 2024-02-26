// Code generated by ent, DO NOT EDIT.

package devicetoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldID, id))
}

// DeviceCode applies equality check predicate on the "device_code" field. It's identical to DeviceCodeEQ.
func DeviceCode(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldDeviceCode, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldStatus, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldToken, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldExpiry, v))
}

// LastRequest applies equality check predicate on the "last_request" field. It's identical to LastRequestEQ.
func LastRequest(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldLastRequest, v))
}

// PollInterval applies equality check predicate on the "poll_interval" field. It's identical to PollIntervalEQ.
func PollInterval(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldPollInterval, v))
}

// CodeChallenge applies equality check predicate on the "code_challenge" field. It's identical to CodeChallengeEQ.
func CodeChallenge(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldCodeChallenge, v))
}

// CodeChallengeMethod applies equality check predicate on the "code_challenge_method" field. It's identical to CodeChallengeMethodEQ.
func CodeChallengeMethod(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldCodeChallengeMethod, v))
}

// DeviceCodeEQ applies the EQ predicate on the "device_code" field.
func DeviceCodeEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldDeviceCode, v))
}

// DeviceCodeNEQ applies the NEQ predicate on the "device_code" field.
func DeviceCodeNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldDeviceCode, v))
}

// DeviceCodeIn applies the In predicate on the "device_code" field.
func DeviceCodeIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldDeviceCode, vs...))
}

// DeviceCodeNotIn applies the NotIn predicate on the "device_code" field.
func DeviceCodeNotIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldDeviceCode, vs...))
}

// DeviceCodeGT applies the GT predicate on the "device_code" field.
func DeviceCodeGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldDeviceCode, v))
}

// DeviceCodeGTE applies the GTE predicate on the "device_code" field.
func DeviceCodeGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldDeviceCode, v))
}

// DeviceCodeLT applies the LT predicate on the "device_code" field.
func DeviceCodeLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldDeviceCode, v))
}

// DeviceCodeLTE applies the LTE predicate on the "device_code" field.
func DeviceCodeLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldDeviceCode, v))
}

// DeviceCodeContains applies the Contains predicate on the "device_code" field.
func DeviceCodeContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContains(FieldDeviceCode, v))
}

// DeviceCodeHasPrefix applies the HasPrefix predicate on the "device_code" field.
func DeviceCodeHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasPrefix(FieldDeviceCode, v))
}

// DeviceCodeHasSuffix applies the HasSuffix predicate on the "device_code" field.
func DeviceCodeHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasSuffix(FieldDeviceCode, v))
}

// DeviceCodeEqualFold applies the EqualFold predicate on the "device_code" field.
func DeviceCodeEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEqualFold(FieldDeviceCode, v))
}

// DeviceCodeContainsFold applies the ContainsFold predicate on the "device_code" field.
func DeviceCodeContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContainsFold(FieldDeviceCode, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContainsFold(FieldStatus, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...[]byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...[]byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldToken, v))
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotNull(FieldToken))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldExpiry, v))
}

// LastRequestEQ applies the EQ predicate on the "last_request" field.
func LastRequestEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldLastRequest, v))
}

// LastRequestNEQ applies the NEQ predicate on the "last_request" field.
func LastRequestNEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldLastRequest, v))
}

// LastRequestIn applies the In predicate on the "last_request" field.
func LastRequestIn(vs ...time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldLastRequest, vs...))
}

// LastRequestNotIn applies the NotIn predicate on the "last_request" field.
func LastRequestNotIn(vs ...time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldLastRequest, vs...))
}

// LastRequestGT applies the GT predicate on the "last_request" field.
func LastRequestGT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldLastRequest, v))
}

// LastRequestGTE applies the GTE predicate on the "last_request" field.
func LastRequestGTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldLastRequest, v))
}

// LastRequestLT applies the LT predicate on the "last_request" field.
func LastRequestLT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldLastRequest, v))
}

// LastRequestLTE applies the LTE predicate on the "last_request" field.
func LastRequestLTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldLastRequest, v))
}

// PollIntervalEQ applies the EQ predicate on the "poll_interval" field.
func PollIntervalEQ(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldPollInterval, v))
}

// PollIntervalNEQ applies the NEQ predicate on the "poll_interval" field.
func PollIntervalNEQ(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldPollInterval, v))
}

// PollIntervalIn applies the In predicate on the "poll_interval" field.
func PollIntervalIn(vs ...int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldPollInterval, vs...))
}

// PollIntervalNotIn applies the NotIn predicate on the "poll_interval" field.
func PollIntervalNotIn(vs ...int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldPollInterval, vs...))
}

// PollIntervalGT applies the GT predicate on the "poll_interval" field.
func PollIntervalGT(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldPollInterval, v))
}

// PollIntervalGTE applies the GTE predicate on the "poll_interval" field.
func PollIntervalGTE(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldPollInterval, v))
}

// PollIntervalLT applies the LT predicate on the "poll_interval" field.
func PollIntervalLT(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldPollInterval, v))
}

// PollIntervalLTE applies the LTE predicate on the "poll_interval" field.
func PollIntervalLTE(v int) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldPollInterval, v))
}

// CodeChallengeEQ applies the EQ predicate on the "code_challenge" field.
func CodeChallengeEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldCodeChallenge, v))
}

// CodeChallengeNEQ applies the NEQ predicate on the "code_challenge" field.
func CodeChallengeNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldCodeChallenge, v))
}

// CodeChallengeIn applies the In predicate on the "code_challenge" field.
func CodeChallengeIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldCodeChallenge, vs...))
}

// CodeChallengeNotIn applies the NotIn predicate on the "code_challenge" field.
func CodeChallengeNotIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldCodeChallenge, vs...))
}

// CodeChallengeGT applies the GT predicate on the "code_challenge" field.
func CodeChallengeGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldCodeChallenge, v))
}

// CodeChallengeGTE applies the GTE predicate on the "code_challenge" field.
func CodeChallengeGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldCodeChallenge, v))
}

// CodeChallengeLT applies the LT predicate on the "code_challenge" field.
func CodeChallengeLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldCodeChallenge, v))
}

// CodeChallengeLTE applies the LTE predicate on the "code_challenge" field.
func CodeChallengeLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldCodeChallenge, v))
}

// CodeChallengeContains applies the Contains predicate on the "code_challenge" field.
func CodeChallengeContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContains(FieldCodeChallenge, v))
}

// CodeChallengeHasPrefix applies the HasPrefix predicate on the "code_challenge" field.
func CodeChallengeHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasPrefix(FieldCodeChallenge, v))
}

// CodeChallengeHasSuffix applies the HasSuffix predicate on the "code_challenge" field.
func CodeChallengeHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasSuffix(FieldCodeChallenge, v))
}

// CodeChallengeEqualFold applies the EqualFold predicate on the "code_challenge" field.
func CodeChallengeEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEqualFold(FieldCodeChallenge, v))
}

// CodeChallengeContainsFold applies the ContainsFold predicate on the "code_challenge" field.
func CodeChallengeContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContainsFold(FieldCodeChallenge, v))
}

// CodeChallengeMethodEQ applies the EQ predicate on the "code_challenge_method" field.
func CodeChallengeMethodEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEQ(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodNEQ applies the NEQ predicate on the "code_challenge_method" field.
func CodeChallengeMethodNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNEQ(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodIn applies the In predicate on the "code_challenge_method" field.
func CodeChallengeMethodIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldIn(FieldCodeChallengeMethod, vs...))
}

// CodeChallengeMethodNotIn applies the NotIn predicate on the "code_challenge_method" field.
func CodeChallengeMethodNotIn(vs ...string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldNotIn(FieldCodeChallengeMethod, vs...))
}

// CodeChallengeMethodGT applies the GT predicate on the "code_challenge_method" field.
func CodeChallengeMethodGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGT(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodGTE applies the GTE predicate on the "code_challenge_method" field.
func CodeChallengeMethodGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldGTE(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodLT applies the LT predicate on the "code_challenge_method" field.
func CodeChallengeMethodLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLT(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodLTE applies the LTE predicate on the "code_challenge_method" field.
func CodeChallengeMethodLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldLTE(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodContains applies the Contains predicate on the "code_challenge_method" field.
func CodeChallengeMethodContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContains(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodHasPrefix applies the HasPrefix predicate on the "code_challenge_method" field.
func CodeChallengeMethodHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasPrefix(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodHasSuffix applies the HasSuffix predicate on the "code_challenge_method" field.
func CodeChallengeMethodHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldHasSuffix(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodEqualFold applies the EqualFold predicate on the "code_challenge_method" field.
func CodeChallengeMethodEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldEqualFold(FieldCodeChallengeMethod, v))
}

// CodeChallengeMethodContainsFold applies the ContainsFold predicate on the "code_challenge_method" field.
func CodeChallengeMethodContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(sql.FieldContainsFold(FieldCodeChallengeMethod, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceToken) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceToken) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceToken) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
