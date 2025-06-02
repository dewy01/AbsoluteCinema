import { createContext, useContext, useState } from 'react';

interface AuthContextProps {
  isAuthenticated: boolean;
  isAdmin: boolean;
  login: (accessToken: string, refreshToken: string, userRole: string) => void;
  logout: () => void;
}

const authContext = createContext<AuthContextProps | undefined>(undefined);

type Props = { children: React.ReactNode };

export const AuthProvider = ({ children }: Props) => {
  const [isAuthenticated, setIsAuthenticated] = useState(() => {
    return (
      localStorage.getItem('accessToken') !== null &&
      localStorage.getItem('refreshToken') !== null &&
      localStorage.getItem('userRole') !== null
    );
  });

  const [isAdmin, setIsAdmin] = useState(() => {
    return localStorage.getItem('userRole') === 'admin';
  });

  const login = (accessToken: string, refreshToken: string, userRole: string) => {
    localStorage.setItem('accessToken', accessToken);
    localStorage.setItem('refreshToken', refreshToken);
    localStorage.setItem('userRole', userRole);
    setIsAuthenticated(true);
    setIsAdmin(userRole === 'admin');
  };

  const logout = () => {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    localStorage.removeItem('userRole');
    setIsAuthenticated(false);
    setIsAdmin(false);
  };

  return (
    <authContext.Provider value={{ isAuthenticated, isAdmin, login, logout }}>
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
