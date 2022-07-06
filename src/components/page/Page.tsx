import { Card, CardBody } from '@patternfly/react-core';
import React from 'react';

import { IPluginPageProps, PageContentSection, PageHeaderSection, PluginPageTitle } from '@kobsio/shared';
import { defaultDescription } from '../../utils/constants';

// The Page component is shown when a user clicks on the plugin instance on the plugins page. You can return whatever
// component you want, but for a unified styling across kobs we recommend to use the PageHeaderSection and
// PageContentSection components on each route of your plugin.
//
// Your plugin can also use multiple routes via React Router. For example the Jaeger plugin supports multiple routes:
// https://github.com/kobsio/kobs/blob/d8c48971826b028647dbb57b09b42a82e02fcd74/plugins/plugin-jaeger/src/components/page/Page.tsx#L10
// To reduce the loading times for use and by splitting your frontend code in several files, it would be nice if you
// import and use your different pages like follows:
//   const Trace = lazy(() => import('./Trace'));
//   const Traces = lazy(() => import('./Traces'));
//
//   const Page: React.FunctionComponent<IPluginPageProps> = ({ instance }: IPluginPageProps) => {
//     return (
//       <Suspense
//         fallback={<Spinner style={{ left: '50%', position: 'fixed', top: '50%', transform: 'translate(-50%, -50%)' }} />}
//       >
//         <Routes>
//           <Route path="/" element={<Traces instance={instance} />} />
//           <Route path="/trace/" element={<Trace instance={instance} />} />
//           <Route path="/trace/:traceID" element={<Trace instance={instance} />} />
//         </Routes>
//       </Suspense>
//     );
//   };
const Page: React.FunctionComponent<IPluginPageProps> = ({ instance }: IPluginPageProps) => {
  return (
    <React.Fragment>
      <PageHeaderSection
        component={
          <PluginPageTitle
            satellite={instance.satellite}
            name={instance.name}
            description={instance.description || defaultDescription}
          />
        }
      />

      <PageContentSection hasPadding={true} hasDivider={true} toolbarContent={undefined} panelContent={undefined}>
        <Card isCompact={true}>
          <CardBody>{JSON.stringify(instance)}</CardBody>
        </Card>
      </PageContentSection>
    </React.Fragment>
  );
};

export default Page;
