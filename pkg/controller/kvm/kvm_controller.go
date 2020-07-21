package kvm

import (
	"context"
	// "strconv"

	kvmv1alpha1 "github.com/raju140/kvm-operator/pkg/apis/kvm/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_kvm")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Kvm Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKvm{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("kvm-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Kvm
	err = c.Watch(&source.Kind{Type: &kvmv1alpha1.Kvm{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Kvm
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &kvmv1alpha1.Kvm{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileKvm implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileKvm{}

// ReconcileKvm reconciles a Kvm object
type ReconcileKvm struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Kvm object and makes changes based on the state read
// and what is in the Kvm.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKvm) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Kvm")
	// Fetch the Kvm instance
	instance := &kvmv1alpha1.Kvm{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	reqLogger.Info("error accoured", instance.Spec)
	var ConnectionType string
	var Host string
	if len(instance.Spec.Connection)!=0{
	    ConnectionType = instance.Spec.Connection
	}else{
		ConnectionType = "tcp"
	}
	if len(instance.Spec.Host)!=0{
	    Host = instance.Spec.Host
	}else{
		Host = "127.0.0.1"
	}
	domainName := request.Name

	conn,err := getConnection(Host,"user",ConnectionType)
	if err!=nil{
		reqLogger.Info("error accoured", "Kvm.error", err, "Kvm.Connection", conn,"kvm.Host",Host)
		//reqLogger.Info("error accoured", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
		return reconcile.Result{}, err
	}
	if (instance.Spec.StatusSpec.Status == "create"){
		kvmDetails,er := ValidateData(*instance,domainName)
		if er!=0{
			return reconcile.Result{},err
		}
		err = create(kvmDetails,*conn)	
		if err!=nil{
			reqLogger.Info("Domain Creation Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{},err
		}
		reqLogger.Info("Domain created",  "KVM.Name", request.Name)

	}
	if (instance.Spec.StatusSpec.Status == "update"){
		kvmDetails,er := ValidateData(*instance,domainName)
		if er!=0{
			return reconcile.Result{},err
		}
		dom,err := getKVMDomainByName(domainName,*conn)
		if err!=nil{
			reqLogger.Info("Domain Updation setVCPU Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		err = DomainSetVcpus(kvmDetails,*dom)
		if err!=nil{
			reqLogger.Info("set domin vcpu value", err)
		}
		err = DomainSetMemory(kvmDetails,*dom)

		if err!=nil{
			reqLogger.Info("Domain Updation setMemory Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("Domain Updated",  "KVM.Name", request.Name)
		

	}
	if ( instance.Spec.StatusSpec.Status == "reboot"){
		dom,err := getKVMDomainByName(domainName,*conn)
		err = DomainShutdownReboot(*dom)
		if err!=nil{
			reqLogger.Info("VM ShutdownReboot Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("VM ShutdownReboot",  "KVM.Name", request.Name)
		
	}
	if (instance.Spec.StatusSpec.Status == "shutdown"){
		dom,err := getKVMDomainByName(domainName,*conn)
		err = KvmShutdownDomain(*dom)
		if err!=nil{
			reqLogger.Info("VM Shutdown Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("VM Shutdown",  "KVM.Name", request.Name)
	}
	if (instance.Spec.StatusSpec.Status == "save"){
		dom,err := getKVMDomainByName(domainName,*conn)
		err = SaveDomain(domainName,*dom)
		if err!=nil{
			reqLogger.Info("VM Save Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("VM Saved",  "KVM.Name", request.Name)
		
	}
	if (instance.Spec.StatusSpec.Status == "restore"){
		dom,err := getKVMDomainByName(domainName,*conn)
		err = RestoreDomain(domainName,*dom,*conn)
		if err!=nil{
			reqLogger.Info("Restored Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("Restored",  "KVM.Name", request.Name)
		
	}
	if (instance.Spec.StatusSpec.Status == "destroy"){
		dom,err := getKVMDomainByName(domainName,*conn)
		err = KvmDestroyDomain(*dom)
		if err!=nil{
			reqLogger.Info("Destroyed Failed", "KVM.error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("Distroyed",  "KVM.Name", request.Name)
		
	}
	if (instance.Spec.StatusSpec.Status == "start"){
		dom,err := getKVMDomainByName(domainName,*conn)
		err = DomainAutoStart(*dom)
		if err!=nil{
			reqLogger.Info("error accoured", "KVM.Error", err, "KVM.Name", request.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("Started",  "KVM.Name", request.Name)
	}

	if err != nil{
		reqLogger.Info("error accoured", "KVM.Error", err, "KVM.Name", request.Name)
		return reconcile.Result{}, err
	}
//	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}
