package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalDomainFederation 
type ExternalDomainFederation struct {
    IdentitySource
}
// NewExternalDomainFederation instantiates a new ExternalDomainFederation and sets the default values.
func NewExternalDomainFederation()(*ExternalDomainFederation) {
    m := &ExternalDomainFederation{
        IdentitySource: *NewIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.externalDomainFederation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExternalDomainFederationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalDomainFederationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalDomainFederation(), nil
}
// GetDisplayName gets the displayName property value. The name of the identity source, typically also the domain name. Read only.
func (m *ExternalDomainFederation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainName gets the domainName property value. The domain name. Read only.
func (m *ExternalDomainFederation) GetDomainName()(*string) {
    val, err := m.GetBackingStore().Get("domainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalDomainFederation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySource.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["issuerUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerUri(val)
        }
        return nil
    }
    return res
}
// GetIssuerUri gets the issuerUri property value. The issuerURI of the incoming federation. Read only.
func (m *ExternalDomainFederation) GetIssuerUri()(*string) {
    val, err := m.GetBackingStore().Get("issuerUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExternalDomainFederation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerUri", m.GetIssuerUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the identity source, typically also the domain name. Read only.
func (m *ExternalDomainFederation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainName sets the domainName property value. The domain name. Read only.
func (m *ExternalDomainFederation) SetDomainName(value *string)() {
    err := m.GetBackingStore().Set("domainName", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerUri sets the issuerUri property value. The issuerURI of the incoming federation. Read only.
func (m *ExternalDomainFederation) SetIssuerUri(value *string)() {
    err := m.GetBackingStore().Set("issuerUri", value)
    if err != nil {
        panic(err)
    }
}
// ExternalDomainFederationable 
type ExternalDomainFederationable interface {
    IdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetDomainName()(*string)
    GetIssuerUri()(*string)
    SetDisplayName(value *string)()
    SetDomainName(value *string)()
    SetIssuerUri(value *string)()
}