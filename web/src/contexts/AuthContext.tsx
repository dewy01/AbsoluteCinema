import { createContext, useContext, useState } from 'react';

interface AuthContextProps {
  isAuthenticated: boolean;
  isAdmin: boolean;
  login: (accessToken: string, refreshToken: string) => void;
  logout: () => void;
  checkAdmin: () => void;
}

const authContext = createContext<AuthContextProps | undefined>(undefined);

type Props = { children: React.ReactNode };

export const AuthProvider = ({ children }: Props) => {
  const [isAuthenticated, setIsAuthenticated] = useState(() => {
    return (
      localStorage.getItem('accessToken') !== null && localStorage.getItem('refreshToken') !== null
    );
  });

  const [isAdmin, setIsAdmin] = useState(() => {
    const role = localStorage.getItem('userRole');
    return role === 'admin';
  });

  const login = (accessToken: string, refreshToken: string) => {
    localStorage.setItem('accessToken', accessToken);
    localStorage.setItem('refreshToken', refreshToken);
    setIsAuthenticated(true);
  };

  const logout = () => {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    setIsAuthenticated(false);
  };

  const checkAdmin = () => {
    const role = localStorage.getItem('userRole');
    if (role === 'admin') {
      setIsAdmin(true);
    } else {
      setIsAdmin(false);
    }
  };

  return (
    <authContext.Provider value={{ isAuthenticated, isAdmin, login, logout, checkAdmin }}>
      {children}
    </authContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(authContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
