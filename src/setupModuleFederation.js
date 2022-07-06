/* eslint-disable */
const deps = require('../package.json').dependencies;

// This file is required to mount the plugin components "Instance", "Panel" and "Page" in the React UI of kobs. You have
// to adjust the "name" property for your plugin, which should have the same value as the "PluginType" constant in the
// "cmd/helloworld.go" file.
module.exports = {
  name: 'helloworld',
  filename: 'remoteEntry.js',
  remotes: {},
  exposes: {
    './Instance': './src/components/instance/Instance.tsx',
    './Panel': './src/components/panel/Panel.tsx',
    './Page': './src/components/page/Page.tsx',
  },
  shared: {
    ...deps,
    react: {
      singleton: true,
      requiredVersion: deps.react,
    },
    'react-dom': {
      singleton: true,
      requiredVersion: deps['react-dom'],
    },
    'react-router-dom': {
      singleton: true,
      requiredVersion: deps['react-router-dom'],
    },
    'react-query': {
      singleton: true,
      requiredVersion: deps['react-query'],
    },
    '@patternfly/patternfly': {
      singleton: true,
      requiredVersion: deps['@patternfly/patternfly'],
    },
    '@patternfly/react-core': {
      singleton: true,
      requiredVersion: deps['@patternfly/react-core'],
    },
    '@patternfly/react-icons': {
      singleton: true,
      requiredVersion: deps['@patternfly/react-icons'],
    },
    '@patternfly/react-table': {
      singleton: true,
      requiredVersion: deps['@patternfly/react-table'],
    },
  },
}
/* eslint-enable */
