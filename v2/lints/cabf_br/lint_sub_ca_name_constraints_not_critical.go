package cabf_br

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************
CA Brower Forum Baseline Requirements, Section 7.1.2.2:

   f. nameConstraints (optional)
If present, this extension SHOULD be marked critical*.

* Non-critical Name Constraints are an exception to RFC 5280 (4.2.1.10), however, they MAY be used until the
Name Constraints extension is supported by Application Software Suppliers whose software is used by a
substantial portion of Relying Parties worldwide
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/teamnsrg/zlint/v2/lint"
	"github.com/teamnsrg/zlint/v2/util"
)

type SubCANameConstraintsNotCritical struct{}

func (l *SubCANameConstraintsNotCritical) Initialize() error {
	return nil
}

func (l *SubCANameConstraintsNotCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsSubCA(cert) && util.IsExtInCert(cert, util.NameConstOID)
}

func (l *SubCANameConstraintsNotCritical) Execute(cert *x509.Certificate) *lint.LintResult {
	if ski := util.GetExtFromCert(cert, util.NameConstOID); ski.Critical {
		return &lint.LintResult{Status: lint.Pass}
	} else {
		return &lint.LintResult{Status: lint.Warn}
	}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_sub_ca_name_constraints_not_critical",
		Description:   "Subordinate CA Certificate: NameConstraints if present, SHOULD be marked critical.",
		Citation:      "BRs: 7.1.2.2",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV102Date,
		Lint:          &SubCANameConstraintsNotCritical{},
	})
}
