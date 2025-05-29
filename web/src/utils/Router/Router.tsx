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

const Register = lazy(() =>
  import('@/views/RegisterView').then((module) => ({
    default: module.RegisterView
  }))
);

export const Router = () => {
  return (
    <Suspense fallback={<Progress />}>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Login />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
        </Route>
      </Routes>
    </Suspense>
  );
};
