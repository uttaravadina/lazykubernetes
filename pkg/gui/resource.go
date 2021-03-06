package gui

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/yolossn/lazykubernetes/pkg/utils"
	duration "k8s.io/apimachinery/pkg/util/duration"

	"github.com/jesseduffield/gocui"
	"sigs.k8s.io/yaml"
)

func (gui *Gui) getResourceView() *gocui.View {
	v, _ := gui.g.View("resource")
	return v
}

func (gui *Gui) onResourceClick(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	infoView := gui.getInfoView()
	// render pod
	switch getResourceTabs()[gui.panelStates.Resource.TabIndex] {
	case "pod":
		infoView.Tabs = getPodInfoTabs()
		gui.panelStates.Resource.SelectedLine = gui.FindSelectedLine(v, len(gui.data.PodData))
		return gui.handlePodSelect(v)
	case "job":
		infoView.Tabs = getJobInfoTabs()
		gui.panelStates.Resource.SelectedLine = gui.FindSelectedLine(v, len(gui.data.JobData))
		return gui.handleJobSelect(v)
	case "deploy":
		infoView.Tabs = getDeployInfoTabs()
		gui.panelStates.Resource.SelectedLine = gui.FindSelectedLine(v, len(gui.data.DeploymentData))
		return gui.handleDeploymentSelect(v)
	case "service":
		infoView.Tabs = getServiceInfoTabs()
		gui.panelStates.Resource.SelectedLine = gui.FindSelectedLine(v, len(gui.data.ServiceData))
		return gui.handleServiceSelect(v)
	case "configMap":
		infoView.Tabs = getConfigMapInfoTabs()
		gui.panelStates.Resource.SelectedLine = gui.FindSelectedLine(v, len(gui.data.ConfigMapData))
		return gui.handleConfigMapSelect(v)
	case "secret":
		infoView.Tabs = getSecretInfoTabs()
		gui.panelStates.Resource.SelectedLine = gui.FindSelectedLine(v, len(gui.data.SecretData))
		return gui.handleSecretSelect(v)
	}

	return nil
}

func (gui *Gui) handleJobSelect(v *gocui.View) error {
	// Find Selected Job
	jobSelected := gui.panelStates.Resource.SelectedLine

	if jobSelected < 0 || jobSelected > len(gui.data.JobData) {
		return nil
	}

	job := gui.data.JobData[jobSelected]

	infoView := gui.getInfoView()

	err := gui.focusPoint(0, gui.panelStates.Resource.SelectedLine, len(gui.data.JobData), v)
	if err != nil {
		return err
	}

	// Find the tab in info panel
	switch getJobInfoTabs()[gui.panelStates.Info.TabIndex] {
	case "description":
		infoView.Clear()
		data, err := gui.k8sClient.DescribeJob(job.Namespace, job.Name)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Fprintln(infoView, string(output))
	}
	return nil
}

func (gui *Gui) handleDeploymentSelect(v *gocui.View) error {
	// Find Selected Deployment
	deploymentSelected := gui.panelStates.Resource.SelectedLine

	if deploymentSelected < 0 || deploymentSelected > len(gui.data.DeploymentData) {
		return nil
	}

	deployment := gui.data.DeploymentData[deploymentSelected]

	infoView := gui.getInfoView()

	err := gui.focusPoint(0, gui.panelStates.Resource.SelectedLine, len(gui.data.DeploymentData), v)
	if err != nil {
		return err
	}

	// Find the tab in info panel
	switch getDeployInfoTabs()[gui.panelStates.Info.TabIndex] {
	case "description":
		infoView.Clear()
		data, err := gui.k8sClient.DescribeDeployment(deployment.Namespace, deployment.Name)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Fprintln(infoView, string(output))
	}
	return nil
}

func (gui *Gui) handleConfigMapSelect(v *gocui.View) error {
	// Find Selected ConfigMap
	configMapSelected := gui.panelStates.Resource.SelectedLine

	if configMapSelected < 0 || configMapSelected > len(gui.data.ConfigMapData) {
		return nil
	}

	configMap := gui.data.ConfigMapData[configMapSelected]

	infoView := gui.getInfoView()

	err := gui.focusPoint(0, gui.panelStates.Resource.SelectedLine, len(gui.data.ConfigMapData), v)
	if err != nil {
		return err
	}

	// Find the tab in info panel
	switch getConfigMapInfoTabs()[gui.panelStates.Info.TabIndex] {
	case "description":
		infoView.Clear()
		data, err := gui.k8sClient.DescribeConfigMap(configMap.Namespace, configMap.Name)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Fprintln(infoView, string(output))
	}
	return nil
}

func (gui *Gui) handleSecretSelect(v *gocui.View) error {
	// Find Selected Secret
	secretSelected := gui.panelStates.Resource.SelectedLine

	if secretSelected < 0 || secretSelected > len(gui.data.SecretData) {
		return nil
	}

	secret := gui.data.SecretData[secretSelected]

	infoView := gui.getInfoView()

	err := gui.focusPoint(0, gui.panelStates.Resource.SelectedLine, len(gui.data.SecretData), v)
	if err != nil {
		return err
	}

	// Find the tab in info panel
	switch getSecretInfoTabs()[gui.panelStates.Info.TabIndex] {
	case "description":
		infoView.Clear()
		data, err := gui.k8sClient.DescribeSecret(secret.Namespace, secret.Name)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Fprintln(infoView, string(output))
	}
	return nil
}

func (gui *Gui) handleServiceSelect(v *gocui.View) error {
	// Find Selected Service
	serviceSelected := gui.panelStates.Resource.SelectedLine

	if serviceSelected < 0 || serviceSelected > len(gui.data.ServiceData) {
		return nil
	}

	service := gui.data.ServiceData[serviceSelected]

	infoView := gui.getInfoView()

	err := gui.focusPoint(0, gui.panelStates.Resource.SelectedLine, len(gui.data.ServiceData), v)
	if err != nil {
		return err
	}

	// Find the tab in info panel
	switch getServiceInfoTabs()[gui.panelStates.Info.TabIndex] {
	case "description":
		infoView.Clear()
		data, err := gui.k8sClient.DescribeService(service.Namespace, service.Name)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Fprintln(infoView, string(output))
	}
	return nil
}

func (gui *Gui) handlePodSelect(v *gocui.View) error {

	// Find Selected Pod
	podSelected := gui.panelStates.Resource.SelectedLine
	if podSelected < 0 || podSelected > len(gui.data.PodData) {
		return nil
	}

	pod := gui.data.PodData[podSelected]

	infoView := gui.getInfoView()

	err := gui.focusPoint(0, gui.panelStates.Resource.SelectedLine, len(gui.data.PodData), v)
	if err != nil {
		return err
	}

	// Find the tab in info panel
	switch getPodInfoTabs()[gui.panelStates.Info.TabIndex] {
	case "logs":
		infoView.Clear()
		gui.g.Update(func(*gocui.Gui) error {
			ctx := context.TODO()
			req := gui.k8sClient.StreamPodLogs(pod.Namespace, pod.Name)
			readCloser, err := req.Stream(ctx)
			if err != nil {
				fmt.Println(err)
			}
			infoView.Clear()
			infoView.Autoscroll = true
			go func() {
				for {
					io.Copy(infoView, readCloser)
				}
			}()
			return nil
		})
	case "description":
		infoView.Clear()
		data, err := gui.k8sClient.DescribePod(pod.Namespace, pod.Name)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(data)
		if err != nil {
			return err
		}

		fmt.Fprintln(infoView, string(output))
	}
	return nil
}

func (gui *Gui) onResourceTabClick(tabIndex int) error {

	resourceView := gui.getResourceView()
	resourceView.TabIndex = tabIndex

	gui.panelStates.Resource.TabIndex = tabIndex
	infoView := gui.getInfoView()
	switch gui.getCurrentResourceTab() {
	case "pod":
		infoView.Tabs = getPodInfoTabs()
	case "job":
		infoView.Tabs = getJobInfoTabs()
	case "deploy":
		infoView.Tabs = getDeployInfoTabs()
	case "service":
		infoView.Tabs = getServiceInfoTabs()
	case "secret":
		infoView.Tabs = getSecretInfoTabs()
	case "configMap":
		infoView.Tabs = getConfigMapInfoTabs()
	}

	return gui.reRenderResource()
}

func (gui *Gui) reRenderResource() error {
	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	ns := gui.getCurrentNS()
	switch getResourceTabs()[gui.panelStates.Resource.TabIndex] {
	case "pod":
		gui.setPods(ns)
		return gui.renderPods()
	case "job":
		gui.setJobs(ns)
		return gui.renderJobs()
	case "deploy":
		gui.setDeployments(ns)
		return gui.renderDeployments()
	case "service":
		gui.setServices(ns)
		return gui.renderServices()
	case "secret":
		gui.setSecrets(ns)
		return gui.renderSecrets()
	case "configMap":
		gui.setConfigMaps(ns)
		return gui.renderConfigMaps()
	}

	return nil
}

func (gui *Gui) getCurrentResourceTab() string {
	return getResourceTabs()[gui.panelStates.Resource.TabIndex]
}

func (gui *Gui) setServices(namespace string) {
	gui.data.rsMux.Lock()
	defer gui.data.rsMux.Unlock()

	services, err := gui.k8sClient.ListServices(namespace)
	if err != nil {

	}
	gui.data.ServiceData = services
}

func (gui *Gui) setPods(namespace string) {
	gui.data.rsMux.Lock()
	defer gui.data.rsMux.Unlock()

	pods, err := gui.k8sClient.ListPods(namespace)
	if err != nil {
		return
	}
	gui.data.PodData = pods
}

func (gui *Gui) setJobs(namespace string) {
	gui.data.rsMux.Lock()
	defer gui.data.rsMux.Unlock()

	jobs, err := gui.k8sClient.ListJobs(namespace)
	if err != nil {
		return
	}
	gui.data.JobData = jobs
}

func (gui *Gui) setDeployments(namespace string) {
	gui.data.rsMux.Lock()
	defer gui.data.rsMux.Unlock()

	deployments, err := gui.k8sClient.ListDeployments(namespace)
	if err != nil {

	}
	gui.data.DeploymentData = deployments
}

func (gui *Gui) setConfigMaps(namespace string) {
	gui.data.rsMux.Lock()
	defer gui.data.rsMux.Unlock()

	configmaps, err := gui.k8sClient.ListConfigMap(namespace)
	if err != nil {
		return
	}
	gui.data.ConfigMapData = configmaps
}

func (gui *Gui) setSecrets(namespace string) {
	gui.data.rsMux.Lock()
	defer gui.data.rsMux.Unlock()

	secrets, err := gui.k8sClient.ListSecrets(namespace)
	if err != nil {
		return
	}
	gui.data.SecretData = secrets
}

func (gui *Gui) renderServices() error {
	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	gui.data.rsMux.RLock()
	defer gui.data.rsMux.RUnlock()

	rsView.Clear()
	services := gui.data.ServiceData
	data := make([][]string, cap(services))
	for i := 0; i < cap(services); i++ {
		data[i] = make([]string, 6)
	}
	// headers := []string{"NAME", "TYPE", "CLUSTER-IP", "EXTERNAL-IP", "PORT(s)", "AGE"}
	headers := []string{"NAME", "TYPE", "CLUSTER-IP", "PORT(s)", "AGE"}

	for i, service := range services {
		data[i][0] = service.Name
		data[i][1] = service.Type
		data[i][2] = service.ClusterIP
		// data[i][3] = service.ExternalIP
		s := ""
		for _, portInfo := range service.Ports {
			port_string := fmt.Sprintf("%v/%s", portInfo.Port, portInfo.Protocol)
			s = fmt.Sprintf("%s %s", s, port_string)
		}
		data[i][3] = fmt.Sprintf("%v", s)
		data[i][4] = duration.HumanDuration(time.Since(service.CreatedAt))
	}

	utils.RenderTable(rsView, data, headers)

	return nil
}

func (gui *Gui) renderSecrets() error {
	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	gui.data.rsMux.RLock()
	defer gui.data.rsMux.RUnlock()

	rsView.Clear()
	secrets := gui.data.SecretData
	data := make([][]string, cap(secrets))

	for i := 0; i < cap(secrets); i++ {
		data[i] = make([]string, 4)
	}
	headers := []string{"NAME", "TYPE", "DATA", "AGE"}

	for i, secret := range secrets {
		data[i][0] = secret.Name
		data[i][1] = secret.Type
		data[i][2] = fmt.Sprintf("%v", secret.Data)
		data[i][3] = duration.HumanDuration(time.Since(secret.CreatedAt))
	}

	utils.RenderTable(rsView, data, headers)

	return nil
}

func (gui *Gui) renderConfigMaps() error {
	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	gui.data.rsMux.RLock()
	defer gui.data.rsMux.RUnlock()

	rsView.Clear()
	configmaps := gui.data.ConfigMapData
	data := make([][]string, cap(configmaps))

	for i := 0; i < cap(configmaps); i++ {
		data[i] = make([]string, 3)
	}
	headers := []string{"NAME", "DATA", "AGE"}

	for i, configmap := range configmaps {
		data[i][0] = configmap.Name
		data[i][1] = fmt.Sprintf("%v", configmap.Data)
		data[i][2] = duration.HumanDuration(time.Since(configmap.CreatedAt))
	}

	utils.RenderTable(rsView, data, headers)

	return nil
}

func (gui *Gui) renderDeployments() error {
	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	gui.data.rsMux.RLock()
	defer gui.data.rsMux.RUnlock()

	rsView.Clear()
	deployments := gui.data.DeploymentData
	data := make([][]string, cap(deployments))

	for i := 0; i < cap(deployments); i++ {
		data[i] = make([]string, 5)
	}
	headers := []string{"NAME", "READY", "UP-TO-DATE", "AVAILABLE", "AGE"}
	for i, deployment := range deployments {
		data[i][0] = deployment.Name
		data[i][1] = fmt.Sprintf("%v/%v", deployment.ReadyReplicas, deployment.Replicas)
		data[i][2] = fmt.Sprintf("%v", deployment.UpdatedReplicas)
		data[i][3] = fmt.Sprintf("%v", deployment.Available)
		data[i][4] = duration.HumanDuration(time.Since(deployment.CreatedAt))
	}

	utils.RenderTable(rsView, data, headers)

	return nil
}

func (gui *Gui) renderJobs() error {
	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	gui.data.rsMux.RLock()
	defer gui.data.rsMux.RUnlock()

	rsView.Clear()
	jobs := gui.data.JobData

	data := make([][]string, cap(jobs))

	for x := 0; x < cap(jobs); x++ {
		data[x] = make([]string, 4)
	}
	headers := []string{"NAME", "COMPLETIONS", "DURATION", "AGE"}
	for i, job := range jobs {
		data[i][0] = job.Name
		data[i][1] = fmt.Sprintf("%v/%v", job.Succeeded, job.Succeeded+job.Failed)
		data[i][2] = duration.HumanDuration(job.CompletedAt.Sub(job.CreatedAt))
		data[i][3] = duration.HumanDuration(time.Since(job.CreatedAt))
	}

	utils.RenderTable(rsView, data, headers)

	return nil
}

func (gui *Gui) renderPods() error {

	rsView := gui.getResourceView()
	if rsView == nil {
		return nil
	}

	gui.data.rsMux.RLock()
	defer gui.data.rsMux.RUnlock()

	rsView.Clear()
	pods := gui.data.PodData

	data := make([][]string, cap(pods))

	for x := 0; x < cap(pods); x++ {
		data[x] = make([]string, 5)
	}
	headers := []string{"NAME", "READY", "STATUS", "RESTARTS", "AGE"}
	for i, pod := range pods {
		data[i][0] = pod.Name
		data[i][1] = fmt.Sprintf("%v/%v", pod.ReadyContainers, pod.TotalContainers)
		data[i][2] = pod.Status
		data[i][3] = fmt.Sprintf("%v", pod.Restarts)
		data[i][4] = duration.HumanDuration(time.Since(pod.CreatedAt))
	}

	utils.RenderTable(rsView, data, headers)

	return nil
}

func (gui *Gui) WatchPods() error {
	_ = gui.reRenderResource()
	// TODO: Handle error
	event, _ := gui.k8sClient.WatchPods("")
	for {
		_ = <-event.ResultChan()
		if gui.getCurrentResourceTab() != "pod" {
			continue
		}
		_ = gui.reRenderNamespace()
	}
}

func (gui *Gui) handleResourceKeyUp(g *gocui.Gui, v *gocui.View) error {
	switch gui.getCurrentResourceTab() {
	case "pod":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.PodData), false)
		return gui.handlePodSelect(v)
	case "job":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.JobData), false)
		return gui.handleJobSelect(v)
	case "deploy":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.DeploymentData), true)
		return gui.handleDeploymentSelect(v)
	case "service":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.ServiceData), false)
		return gui.handleServiceSelect(v)
	case "secret":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.SecretData), false)
		return gui.handleSecretSelect(v)
	case "configMap":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.ConfigMapData), false)
		return gui.handleConfigMapSelect(v)
	}
	return nil
}

func (gui *Gui) handleResourceKeyDown(g *gocui.Gui, v *gocui.View) error {
	switch gui.getCurrentResourceTab() {
	case "pod":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.PodData), true)
		return gui.handlePodSelect(v)
	case "job":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.JobData), true)
		return gui.handleJobSelect(v)
	case "deploy":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.DeploymentData), true)
		return gui.handleDeploymentSelect(v)
	case "service":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.ServiceData), false)
		return gui.handleServiceSelect(v)
	case "secret":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.SecretData), false)
		return gui.handleSecretSelect(v)
	case "configMap":
		gui.changeSelectedLine(&gui.panelStates.Resource.SelectedLine, len(gui.data.ConfigMapData), false)
		return gui.handleConfigMapSelect(v)
	}
	return nil
}
