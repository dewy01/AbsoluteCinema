import Progress from '@/components/Progress/Progress';
import { lazy, Suspense } from 'react';
import { Route, Routes } from 'react-router-dom';
import { AuthGuard } from '../AuthGuard';

const Forbidden = lazy(() =>
  import('@/views/ForbiddenView').then((module) => ({
    default: module.ForbiddenView
  }))
);

const NotFound = lazy(() =>
  import('@/views/NotFoundView').then((module) => ({
    default: module.NotFoundView
  }))
);

const Layout = lazy(() =>
  import('@/views/LayoutView').then((module) => ({
    default: module.LayoutView
  }))
);

const Home = lazy(() =>
  import('@/views/HomeView').then((module) => ({
    default: module.HomeView
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

const Movie = lazy(() =>
  import('@/views/MovieView').then((module) => ({
    default: module.MovieView
  }))
);

const Screening = lazy(() =>
  import('@/views/ScreeningView').then((module) => ({
    default: module.ScreeningView
  }))
);

const Reservation = lazy(() =>
  import('@/views/ReservationView').then((module) => ({
    default: module.ReservationView
  }))
);

const UserReservation = lazy(() =>
  import('@/views/UserReservationsView').then((module) => ({
    default: module.UserReservationsView
  }))
);

const UpdateReservation = lazy(() =>
  import('@/views/ReservationView').then((module) => ({
    default: module.UpdateReservationView
  }))
);

export const Router = () => {
  return (
    <Suspense fallback={<Progress />}>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          <Route path="movie/:id" element={<Movie />} />
          <Route path="screening/:id" element={<Screening />} />
          <Route path="reservation/:id" element={<Reservation />} />
          <Route
            path="my-reservations"
            element={
              <AuthGuard>
                <UserReservation />
              </AuthGuard>
            }
          />
          <Route path="reservation/:id/update" element={<UpdateReservation />} />
          <Route path="not-found" element={<NotFound />} />
          <Route path="forbidden" element={<Forbidden />} />
          <Route path="*" element={<NotFound />} />
        </Route>
      </Routes>
    </Suspense>
  );
};
