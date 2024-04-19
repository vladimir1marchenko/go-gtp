// Copyright 2019-2024 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"

	"github.com/wmnsk/go-gtp/gtpv2"
	"github.com/wmnsk/go-gtp/gtpv2/ie"
	"github.com/wmnsk/go-gtp/gtpv2/message"
	"github.com/wmnsk/go-gtp/gtpv2/testutils"
)

func TestCreateSessionRequest(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/FromMMEtoSGW",
			Structured: message.NewCreateSessionRequest(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ie.NewIMSI("123451234567890"),
				ie.NewMSISDN("123450123456789"),
				ie.NewAccessPointName("some.apn.example"),
				ie.NewFullyQualifiedTEID(gtpv2.IFTypeS11MMEGTPC, 0xffffffff, "1.1.1.1", ""),
				ie.NewFullyQualifiedTEID(gtpv2.IFTypeS5S8PGWGTPC, 0xffffffff, "1.1.1.2", "").WithInstance(1),
				ie.NewPDNType(gtpv2.PDNTypeIPv4),
				ie.NewAggregateMaximumBitRate(0x11111111, 0x22222222),
				ie.NewIndicationFromOctets(0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40),
				ie.NewBearerContext(
					ie.NewEPSBearerID(0x05),
					ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
				),
				ie.NewBearerContext(
					ie.NewEPSBearerID(0x06),
					ie.NewBearerQoS(1, 2, 1, 0xff, 0x1111111111, 0x2222222222, 0x1111111111, 0x2222222222),
				),
				ie.NewMobileEquipmentIdentity("123450123456789"),
				ie.NewServingNetwork("123", "45"),
				ie.NewPDNAddressAllocation("2.2.2.2"),
				ie.NewAPNRestriction(gtpv2.APNRestrictionPublic1),
				ie.NewUserLocationInformationStruct(
					nil, nil, nil, ie.NewTAI("123", "45", 0x0001),
					ie.NewECGI("123", "45", 0x00000101), nil, nil, nil,
				),
				ie.NewRATType(gtpv2.RATTypeEUTRAN),
				ie.NewSelectionMode(gtpv2.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
			),
			Serialized: []byte{
				// Header
				0x48, 0x20, 0x00, 0xed, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// IMSI
				0x01, 0x00, 0x08, 0x00, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0,
				// MSISDN
				0x4c, 0x00, 0x08, 0x00, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9,
				// MEI
				0x4b, 0x00, 0x08, 0x00, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9,
				// ULI: TAI ECGI
				0x56, 0x00, 0x0d, 0x00, 0x18,
				0x21, 0xf3, 0x54, 0x00, 0x01,
				0x21, 0xf3, 0x54, 0x00, 0x00, 0x01, 0x01,
				// ServingNetwork
				0x53, 0x00, 0x03, 0x00, 0x21, 0xf3, 0x54,
				// RATType
				0x52, 0x00, 0x01, 0x00, 0x06,
				// Indication
				0x4d, 0x00, 0x07, 0x00, 0xa1, 0x08, 0x15, 0x10, 0x88, 0x81, 0x40,
				// F-TEID S11
				0x57, 0x00, 0x09, 0x00, 0x8a, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x01,
				// F-TEID S5/S8
				0x57, 0x00, 0x09, 0x01, 0x87, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x02,
				// APN
				0x47, 0x00, 0x11, 0x00, 0x04, 0x73, 0x6f, 0x6d, 0x65, 0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
				// SelectionMode
				0x80, 0x00, 0x01, 0x00, 0x00,
				// PDNType
				0x63, 0x00, 0x01, 0x00, 0x01,
				// PAA
				0x4f, 0x00, 0x05, 0x00, 0x01, 0x02, 0x02, 0x02, 0x02,
				// APNRestriction
				0x7f, 0x00, 0x01, 0x00, 0x01,
				// AMBR
				0x48, 0x00, 0x08, 0x00, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22,
				// BearerContext 1
				0x5d, 0x00, 0x1f, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				//   BearerQoS
				0x50, 0x00, 0x16, 0x00, 0x49, 0xff,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22,
				// BearerContext 2
				0x5d, 0x00, 0x1f, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x06,
				//   BearerQoS
				0x50, 0x00, 0x16, 0x00, 0x49, 0xff,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x22, 0x22, 0x22, 0x22, 0x22,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializable, error) {
		v, err := message.ParseCreateSessionRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
