import { useCurrentUser } from '@/apis/User';
import React, { createContext, useContext, useEffect, useState } from 'react';

interface AuthContextProps {
  isAuthenticated: boolean;
  isAdmin: boolean;
  userProps: {
    id?: string;
    name?: string;
    email?: string;
    role?: string;
  };
  logout: () => void;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

type Props = { children: React.ReactNode };

export const AuthProvider = ({ children }: Props) => {
  const { data, isSuccess, isError } = useCurrentUser();
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);
  const [userProps, setUserProps] = useState({
    id: '',
    name: '',
    email: '',
    role: ''
  });

  useEffect(() => {
    if (isSuccess && data) {
      setUserProps({
        id: data.id ?? '',
        name: data.name ?? '',
        email: data.email ?? '',
        role: data.role ?? ''
      });
      setIsAuthenticated(true);
      setIsAdmin(data.role === 'admin');
    } else if (isError) {
      setIsAuthenticated(false);
      setIsAdmin(false);
      setUserProps({ id: '', name: '', email: '', role: '' });
    }
  }, [isSuccess, isError, data]);

  const logout = () => {
    setIsAuthenticated(false);
    setIsAdmin(false);
    setUserProps({ id: '', name: '', email: '', role: '' });
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, isAdmin, userProps, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error('useAuth must be used within AuthProvider');
  return context;
};
