import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';
import { AuthProvider } from './contexts/AuthContext.tsx';
import { QueryClientProvider } from '@tanstack/react-query';
import { queryClient } from './apis/api.ts';
import { CinemaProvider } from './contexts/CinemaContext.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <CinemaProvider>
          <App />
        </CinemaProvider>
      </AuthProvider>
    </QueryClientProvider>
  </StrictMode>
);
