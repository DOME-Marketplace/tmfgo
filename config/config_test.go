package config

import (
	"path/filepath"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	yamlPath := filepath.Join("..", "tmf.yaml")
	cfg, err := readConfigFile(yamlPath)
	if err != nil {
		t.Fatalf("failed to read config file %s: %v", yamlPath, err)
	}

	if cfg.Environment != lclConfig.Environment {
		t.Errorf("expected Environment %q, got %q", lclConfig.Environment, cfg.Environment)
	}
	if cfg.ProxyEnabled != lclConfig.ProxyEnabled {
		t.Errorf("expected ProxyEnabled %v, got %v", lclConfig.ProxyEnabled, cfg.ProxyEnabled)
	}
	if cfg.ServerOperatorOrganizationIdentifier != lclConfig.ServerOperatorOrganizationIdentifier {
		t.Errorf("expected ServerOperatorOrganizationIdentifier %q, got %q", lclConfig.ServerOperatorOrganizationIdentifier, cfg.ServerOperatorOrganizationIdentifier)
	}
	if cfg.ServerOperatorDid != lclConfig.ServerOperatorDid {
		t.Errorf("expected ServerOperatorDid %q, got %q", lclConfig.ServerOperatorDid, cfg.ServerOperatorDid)
	}
	if cfg.ServerOperatorName != lclConfig.ServerOperatorName {
		t.Errorf("expected ServerOperatorName %q, got %q", lclConfig.ServerOperatorName, cfg.ServerOperatorName)
	}
	if cfg.ServerOperatorCountry != lclConfig.ServerOperatorCountry {
		t.Errorf("expected ServerOperatorCountry %q, got %q", lclConfig.ServerOperatorCountry, cfg.ServerOperatorCountry)
	}
	if cfg.LEARPower.Domain != lclConfig.LEARPower.Domain || cfg.LEARPower.Function != lclConfig.LEARPower.Function || cfg.LEARPower.Type != lclConfig.LEARPower.Type {
		t.Errorf("expected LEARPower %+v, got %+v", lclConfig.LEARPower, cfg.LEARPower)
	}
	if cfg.ProductCreatePower.Domain != lclConfig.ProductCreatePower.Domain || cfg.ProductCreatePower.Function != lclConfig.ProductCreatePower.Function {
		t.Errorf("expected ProductCreatePower %+v, got %+v", lclConfig.ProductCreatePower, cfg.ProductCreatePower)
	}
	if cfg.ProductUpdatePower.Domain != lclConfig.ProductUpdatePower.Domain || cfg.ProductUpdatePower.Function != lclConfig.ProductUpdatePower.Function {
		t.Errorf("expected ProductUpdatePower %+v, got %+v", lclConfig.ProductUpdatePower, cfg.ProductUpdatePower)
	}
	if cfg.ProductDeletePower.Domain != lclConfig.ProductDeletePower.Domain || cfg.ProductDeletePower.Function != lclConfig.ProductDeletePower.Function {
		t.Errorf("expected ProductDeletePower %+v, got %+v", lclConfig.ProductDeletePower, cfg.ProductDeletePower)
	}
	if cfg.PolicyFileName != lclConfig.PolicyFileName {
		t.Errorf("expected PolicyFileName %q, got %q", lclConfig.PolicyFileName, cfg.PolicyFileName)
	}
	if cfg.RemoteTMFServer != lclConfig.RemoteTMFServer {
		t.Errorf("expected RemoteTMFServer %q, got %q", lclConfig.RemoteTMFServer, cfg.RemoteTMFServer)
	}
	if cfg.VerifierServer != lclConfig.VerifierServer {
		t.Errorf("expected VerifierServer %q, got %q", lclConfig.VerifierServer, cfg.VerifierServer)
	}
	if cfg.Dbname != lclConfig.Dbname {
		t.Errorf("expected Dbname %q, got %q", lclConfig.Dbname, cfg.Dbname)
	}
	if cfg.ClonePeriod != lclConfig.ClonePeriod {
		t.Errorf("expected ClonePeriod %v, got %v", lclConfig.ClonePeriod, cfg.ClonePeriod)
	}
	if cfg.Features != lclConfig.Features {
		t.Errorf("expected Features %+v, got %+v", lclConfig.Features, cfg.Features)
	}
}
