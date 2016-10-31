package gslb

type Gslbservice struct {
  Appflowlog string `json:"appflowlog,omitempty"`
  Cip string `json:"cip,omitempty"`
  Cipheader string `json:"cipheader,omitempty"`
  Clttimeout int `json:"clttimeout,omitempty"`
  Cnameentry string `json:"cnameentry,omitempty"`
  Comment string `json:"comment,omitempty"`
  Cookietimeout int `json:"cookietimeout,omitempty"`
  Downstateflush string `json:"downstateflush,omitempty"`
  Gslb string `json:"gslb,omitempty"`
  Gslbsvcstats int `json:"gslbsvcstats,omitempty"`
  Gslbthreshold int `json:"gslbthreshold,omitempty"`
  Hashid int `json:"hashid,omitempty"`
  Healthmonitor string `json:"healthmonitor,omitempty"`
  Ip string `json:"ip,omitempty"`
  Ipaddress string `json:"ipaddress,omitempty"`
  Maxaaausers int `json:"maxaaausers,omitempty"`
  Maxbandwidth int `json:"maxbandwidth,omitempty"`
  Maxclient int `json:"maxclient,omitempty"`
  Monitornamesvc string `json:"monitor_name_svc,omitempty"`
  Monitorstate string `json:"monitor_state,omitempty"`
  Monstate string `json:"monstate,omitempty"`
  Monthreshold int `json:"monthreshold,omitempty"`
  Newname string `json:"newname,omitempty"`
  Port int `json:"port,omitempty"`
  Preferredlocation string `json:"preferredlocation,omitempty"`
  Publicip string `json:"publicip,omitempty"`
  Publicport int `json:"publicport,omitempty"`
  Servername string `json:"servername,omitempty"`
  Servicename string `json:"servicename,omitempty"`
  Servicetype string `json:"servicetype,omitempty"`
  Sitename string `json:"sitename,omitempty"`
  Sitepersistence string `json:"sitepersistence,omitempty"`
  Siteprefix string `json:"siteprefix,omitempty"`
  State string `json:"state,omitempty"`
  Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
  Svrstate string `json:"svrstate,omitempty"`
  Svrtimeout int `json:"svrtimeout,omitempty"`
  Viewip string `json:"viewip,omitempty"`
  Viewname string `json:"viewname,omitempty"`
  Weight int `json:"weight,omitempty"`
}