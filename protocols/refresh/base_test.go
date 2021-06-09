package refresh

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taurusgroup/cmp-ecdsa/pb"
	"github.com/taurusgroup/cmp-ecdsa/pkg/party"
	"github.com/taurusgroup/cmp-ecdsa/pkg/round"
	"google.golang.org/protobuf/proto"
)

type testParty struct {
	idx int
	id  party.ID
	s   *round.Session
	r   round.Round
}

func feedMessages(t *testing.T, parties []*testParty, msgs []*pb.Message) {
	for _, msg := range msgs {
		b, err := proto.Marshal(msg)
		require.NoError(t, err, "failed to marshal message")
		for _, p := range parties {
			var m2 pb.Message
			require.NoError(t, proto.Unmarshal(b, &m2), "failed to unmarshal message")
			if m2.GetFromID() == p.id {
				continue
			}
			if m2.To != "" && m2.GetToID() != p.id {
				continue
			}
			require.NoError(t, p.r.ProcessMessage(&m2))
		}
	}
}

func TestRound(t *testing.T) {
	N := 3

	sessions := round.FakeKeygen(N, N-1)

	s1 := sessions[0]
	parties := make([]*testParty, N)
	for idxJ, j := range s1.PartyIDs {
		r, err := NewRound(sessions[idxJ])
		require.NoError(t, err, "round creation should not result in an error")
		parties[idxJ] = &testParty{
			idx: idxJ,
			id:  j,
			s:   sessions[idxJ],
			r:   r,
		}
	}

	// get the first messages
	msgs1 := make([]*pb.Message, 0, N)
	for _, pj := range parties {
		msgs1New, err := pj.r.GenerateMessages()
		require.NoError(t, err, "round 1 GenerateMessages should not cause an error")
		msgs1 = append(msgs1, msgs1New...)
		newR, err := pj.r.Finalize()
		require.NoError(t, err, "round 1 Finalize should not cause an error")
		pj.r = newR
	}

	fmt.Println("R1 done")

	feedMessages(t, parties, msgs1)

	// get the second set of  messages
	msgs2 := make([]*pb.Message, 0, N*N)
	for _, pj := range parties {
		r, ok := pj.r.(*round2)
		if !ok {
			t.Errorf("not the right round")
		}

		msgs2New, err := r.GenerateMessages()
		if err != nil {
			t.Error(err)
		}
		msgs2 = append(msgs2, msgs2New...)

		newR, err := pj.r.Finalize()
		if err != nil {
			t.Error(err)
		}
		pj.r = newR
	}

	fmt.Println("R2 done")

	feedMessages(t, parties, msgs2)

	// get the second set of  messages
	msgs3 := make([]*pb.Message, 0, N*N)
	for _, pj := range parties {
		r, ok := pj.r.(*round3)
		if !ok {
			t.Errorf("not the right round")
		}

		msgs3New, err := r.GenerateMessages()
		if err != nil {
			t.Error(err)
		}
		msgs3 = append(msgs3, msgs3New...)

		newR, err := pj.r.Finalize()
		if err != nil {
			t.Error(err)
		}
		pj.r = newR
	}

	fmt.Println("R3 done")

	feedMessages(t, parties, msgs3)

	// get the third set of  messages
	msgs4 := make([]*pb.Message, 0, N*N)
	for _, pj := range parties {
		r, ok := pj.r.(*round4)
		if !ok {
			t.Errorf("not the right round")
		}

		msgs4New, err := r.GenerateMessages()
		if err != nil {
			t.Error(err)
		}
		msgs4 = append(msgs4, msgs4New...)

		newR, err := pj.r.Finalize()
		if err != nil {
			t.Error(err)
		}
		pj.r = newR
	}

	//// check rid is the same for all
	var rho []byte
	for idx, p := range parties {
		if idx == 0 {
			rho = p.r.(*output).rho
		} else {
			require.EqualValues(t, rho, p.r.(*output).rho, "rhos should be the same")
			//if !bytes.Equal(rho, p.r.(*output).rho) {
			//	t.Error("rho is different")
			//}
		}
	}

	//// check pub key is the same for all
	origPk := s1.PublicKey
	for _, p := range parties {
		newSession := p.r.(*output).S
		err := newSession.RecomputeSSID()
		assert.NoError(t, err, "failed to recompute ssid")
		//newSecret := newSession.Secret
		newPk := newSession.PublicKey
		//require.NoError(t, err, "new session should yield a public key")
		assert.True(t, newPk.Equal(origPk), "new public key should be the same")
		assert.NoError(t, newSession.Validate(), "new session validation should pass")
	}
}