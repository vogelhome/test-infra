// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package common

// Annotations
const (
	// AnnotationResumeTestrun is the annotation name to trigger resume on the testrun
	AnnotationResumeTestrun = "testmachinery.sapcloud.io/resume"

	// AnnotationCollectTestrun is the annotation to trigger collection and persistence of testrun results
	AnnotationCollectTestrun = "testmachinery.garden.cloud/collect"

	// AnnotationSystemStep is the testflow step annotation to specify that the step is a testmachinery system step.
	// It indicates that it should not be considered as a test and therefore should not count for a test to be failed.
	AnnotationSystemStep = "testmachinery.sapcloud.io/system-step"

	// AnnotationTestDefName is the name of origin TestDefinition.
	AnnotationTestDefName = "testmachinery.sapcloud.io/TestDefinition"

	// AnnotationTestDefID is the unique name of origin TestDefinition in a specific flow and step.
	AnnotationTestDefID = "testmachinery.sapcloud.io/ID"

	// LabelTMDashboardIngress is the label to identify TestMachinery ingress objects.
	LabelTMDashboardIngress = "testmachinery.garden.cloud/tm-dashboard"
)

// Metadata Annotations
const (

	// AnnotationTestrunPurpose is the annotation name to specify a purpose of the testrun
	AnnotationTestrunPurpose = "testmachinery.sapcloud.io/purpose"

	// AnnotationTemplateIDTestrun is the annotation to specify the name of the template the testrun is rendered from
	AnnotationTemplateIDTestrun = "testrunner.testmachinery.gardener.cloud/templateID"

	// AnnotationRetries is the annotation to specify the retry count of the current testrun
	AnnotationRetries = "testrunner.testmachinery.gardener.cloud/retries"

	// AnnotationPreviousAttempt is the testrun id if the previous testrun
	AnnotationPreviousAttempt = "testrunner.testmachinery.gardener.cloud/previous-attempt"

	// AnnotationLandscape is the annotation to specify the landscape this testrun is testing
	AnnotationLandscape = "metadata.testmachinery.gardener.cloud/landscape"

	// AnnotationK8sVersion is the annotation to specify the k8s version the testrun is testing
	AnnotationK8sVersion = "metadata.testmachinery.gardener.cloud/k8sVersion"

	// AnnotationCloudProvider is the annotation to specify the cloudprovider the testrun is testing
	AnnotationCloudProvider = "metadata.testmachinery.gardener.cloud/cloudprovider"

	// AnnotationContainerRuntime is the annotation to specify the container runtime of the shoot nodes the testrun is testing
	AnnotationContainerRuntime = "metadata.testmachinery.gardener.cloud/container-runtime"

	// AnnotationOperatingSystem is the annotation to specify the operating system of the shoot nodes the testrun is testing
	AnnotationOperatingSystem = "metadata.testmachinery.gardener.cloud/operating-system"

	// AnnotationOperatingSystemVersion is the annotation to specify the version of the operating system of the shoot nodes the testrun is testing
	AnnotationOperatingSystemVersion = "metadata.testmachinery.gardener.cloud/operating-system-version"

	// AnnotationRegion is the annotation to specify the region of the shoot the testrun is testing
	AnnotationRegion = "metadata.testmachinery.gardener.cloud/region"

	// AnnotationZone is the annotation to specify the zone of the shoot the testrun is testing
	AnnotationZone = "metadata.testmachinery.gardener.cloud/zone"

	// AnnotationAllowPrivilegedContainers is the annotation describing whether and how created shoots will have allowPrivilegedContainers configured
	AnnotationAllowPrivilegedContainers = "metadata.testmachinery.gardener.cloud/allow-privileged-containers"

	// AnnotationShootAnnotations is the annotation describing which additional shoot annotations were set (as they could impact the shoot behaviour)
	AnnotationShootAnnotations = "metadata.testmachinery.gardener.cloud/shoot-annotations"

	// AnnotationFlavorDescription is the annotation to describe the test flavor of the current run testrun
	AnnotationFlavorDescription = "metadata.testmachinery.gardener.cloud/flavor-description"

	// AnnotationDimension is the annotation to specify the dimension the testrun is testing
	AnnotationDimension = "metadata.testmachinery.gardener.cloud/dimension"

	// AnnotationGroupPurpose is the annotation to describe a run group with an arbitrary string
	AnnotationGroupPurpose = "metadata.testmachinery.gardener.cloud/group-purpose"

	// AnnotationTMDashboardURL is the annotation to describe the URL of the TM Dashboard if present
	AnnotationTMDashboardURL = "metadata.testmachinery.gardener.cloud/tm-dashboard-url"

	// LabelTestrunExecutionGroup is the label to specify the unique name of the run (multiple testruns) this test belongs to.
	// A run represents all tests that are running from one testrunner.
	LabelTestrunExecutionGroup = "testrunner.testmachinery.gardener.cloud/execution-group"
)

// Testrunner Annotations
const (
	// LabelUploadedToGithub is the label to specify whether the testrun result was uploaded to github
	LabelUploadedToGithub = "testrunner.testmachinery.sapcloud.io/uploaded-to-github"

	// images
	DockerImageGardenerApiServer = "eu.gcr.io/gardener-project/gardener/apiserver"

	// Repositories
	TestInfraRepo   = "https://github.com/gardener/test-infra.git"
	GardenSetupRepo = "https://github.com/gardener/garden-setup.git"
	GardenerRepo    = "https://github.com/gardener/gardener.git"

	PatternLatest                 = "latest"
	PatternOneMinorBeforeLatest   = "oneMinorBeforeLatest"
	PatternTwoMinorBeforeLatest   = "twoMinorBeforeLatest"
	PatternThreeMinorBeforeLatest = "threeMinorBeforeLatest"
	PatternFourMinorBeforeLatest  = "fourMinorBeforeLatest"

	// TM Dashboard
	DashboardExecutionGroupParameter = "runID"

	// DashboardPaginationFrom is the name of the http parameter for the pagination from index.
	DashboardPaginationFrom = "from"

	// DashboardPaginationTo is the name of the http parameter for the pagination from index.
	DashboardPaginationTo = "to"
)

var (
	// Default timeout of 4 hours to wait before resuming the testrun
	DefaultPauseTimeout = 14400
)
