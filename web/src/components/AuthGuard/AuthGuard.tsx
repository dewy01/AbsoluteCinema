import { useAuth } from '@/contexts';
import React from 'react';
import { Navigate } from 'react-router-dom';

interface AuthGuardProps {
  children: React.ReactElement;
}

export const AuthGuard = ({ children }: AuthGuardProps) => {
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    return <Navigate to="/forbidden" replace />;
  }

  return children;
};
