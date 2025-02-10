package containerruntime

import (
	"net/url"
	"os"
	"strings"
)

// AutoProxyConfiguration configures the container to use the proxy settings from the host environment
func (c *Container) AutoProxyConfiguration() {
	httpProxy := os.Getenv("HTTP_PROXY")
	httpsProxy := os.Getenv("HTTPS_PROXY")
	noProxy := os.Getenv("NO_PROXY")
	c.ProxyConfiguration(httpProxy, httpsProxy, noProxy)
}

// ProxyConfiguration configures the container to use the given proxy settings
func (c *Container) ProxyConfiguration(httpProxy string, httpsProxy string, noProxy string) {
	if httpProxy != "" {
		c.AddEnvironmentVariable("HTTP_PROXY", httpProxy)
		c.AddEnvironmentVariable("http_proxy", httpProxy)
	}
	if httpsProxy != "" {
		c.AddEnvironmentVariable("HTTPS_PROXY", httpsProxy)
		c.AddEnvironmentVariable("https_proxy", httpsProxy)
	}
	if noProxy != "" {
		c.AddEnvironmentVariable("NO_PROXY", noProxy)
		c.AddEnvironmentVariable("no_proxy", noProxy)
	}

	// jvm
	var javaProxyOpts []string
	if httpProxy != "" {
		proxyURL, err := url.Parse(httpProxy)
		if err == nil {
			javaProxyOpts = append(javaProxyOpts, "-Dhttp.proxyHost="+proxyURL.Hostname())
			javaProxyOpts = append(javaProxyOpts, "-Dhttp.proxyPort="+proxyURL.Port())
			javaProxyOpts = append(javaProxyOpts, "-Dhttp.nonProxyHosts="+ToJavaNoPoxyFormat(noProxy))
		}
	}
	if httpsProxy != "" {
		proxyURL, err := url.Parse(httpsProxy)
		if err == nil {
			javaProxyOpts = append(javaProxyOpts, "-Dhttps.proxyHost="+proxyURL.Hostname())
			javaProxyOpts = append(javaProxyOpts, "-Dhttps.proxyPort="+proxyURL.Port())
			javaProxyOpts = append(javaProxyOpts, "-Dhttps.nonProxyHosts="+ToJavaNoPoxyFormat(noProxy))
		}
	}
	if len(javaProxyOpts) > 0 {
		c.AddEnvironmentVariable("CID_PROXY_JVM", strings.Join(javaProxyOpts, " "))
	}
}

func ToJavaNoPoxyFormat(input string) string {
	return strings.ReplaceAll(input, ",", "|")
}
