import Progress from '@/components/Progress/Progress';
import { lazy, Suspense } from 'react';
import { Route, Routes } from 'react-router-dom';

const Layout = lazy(() =>
  import('@/views/LayoutView').then((module) => ({
    default: module.LayoutView
  }))
);
const Login = lazy(() =>
  import('@/views/LoginView').then((module) => ({
    default: module.LoginView
  }))
);

export const Router = () => {
  return (
    <Suspense fallback={<Progress />}>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Login />} />
        </Route>
      </Routes>
    </Suspense>
  );
};
