package pet

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
	"xgmdr.com/pad/internal/api"
	pb "xgmdr.com/pad/proto"
)

var ae = &api.AuthenticatedIdentity{
	Issuer:  "test-issuer",
	Subject: "test-subject",
	Email:   "test@test.com",
}

func TestApiToEntityModelConversion(t *testing.T) {
	npr := &pb.NewPetRequest{
		Name:   "aName",
		Gender: pb.Gender_FEMALE,
		Type:   "Dog",
		Dob:    timestamppb.New(time.Now()),
	}

	pe, _ := toPetEntity(ae, npr)

	if len(pe.Data) == 0 {
		t.Error("no data")
	}

	if pe.Owner.Issuer != ae.Issuer {
		t.Error("non-matching owner issuer")
	}
	if pe.Owner.Email != ae.Email {
		t.Error("non-matching owner email")
	}
	if pe.Owner.OwnerId != ae.Subject {
		t.Error("non-matching owner id")
	}

	fmt.Println(pe.Data)
}
