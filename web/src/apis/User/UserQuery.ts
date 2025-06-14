import { useAuth } from '@/contexts/AuthContext';
import type { components } from '@/types/openapi';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { enqueueSnackbar, useSnackbar } from 'notistack';
import { useNavigate } from 'react-router-dom';
import { queryClient } from '../api';
import {
  deleteUserById,
  getCurrentUser,
  getUserById,
  postUserLogin,
  postUserLogout,
  postUserRegister,
  putUserById
} from './UserApi';

export const callUserRegister = () => {
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  return useMutation({
    mutationKey: ['register'],
    mutationFn: (userData: components['schemas']['CreateUserInput']) => postUserRegister(userData),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Użytkownik został poprawnie zarejestrowany!' });
      navigate('/login', { replace: true });
    },
    onError: (err) => {
      if (err instanceof AxiosError) {
        if (err.response?.status === 500) {
          enqueueSnackbar({
            variant: 'error',
            message: 'Email jest już zajęty.'
          });
        }
      }
    }
  });
};

export const useUserLogin = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['login'],
    mutationFn: postUserLogin,
    onSuccess: async () => {
      queryClient.invalidateQueries({
        queryKey:['me']
      });

      enqueueSnackbar({ message: 'Zalogowano pomyślnie!' });
      navigate('/', { replace: true });
    },
    onError: () => {
      enqueueSnackbar({
        variant: 'error',
        message: 'Nieprawidłowe dane logowania.'
      });
    }
  });
};

export const useUserLogout = () => {
  const navigate = useNavigate();
  const { logout } = useAuth();

  return useMutation({
    mutationKey: ['logout'],
    mutationFn: () => postUserLogout(),
    onSuccess: () => {
      logout();
      navigate('/', { replace: true });
      enqueueSnackbar({ message: 'Wylogowano pomyślnie!' });
      queryClient.invalidateQueries({
        queryKey:['me']
      });
    },
    onError: () => {
      logout();
      navigate('/', { replace: true });
      enqueueSnackbar({ message: 'Wylogowano pomyślnie!' });
    }
  });
};

export const useCurrentUser = () => {
  return useQuery({
    queryKey: ['me'],
    queryFn: getCurrentUser,
  });
};

export const useUserById = (id: string) => {
  return useQuery({
    queryKey: ['user', id],
    queryFn: () => getUserById(id)
  });
};

export const callUpdateUser = (id: string) => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationKey: ['updateUser', id],
    mutationFn: (data: components['schemas']['UpdateUserInput']) => putUserById(id, data),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Użytkownik zaktualizowany!' });
      queryClient.invalidateQueries({ queryKey: ['user', id] });
    },
    onError: () => {
      enqueueSnackbar({ variant: 'error', message: 'Błąd podczas aktualizacji użytkownika.' });
    }
  });
};

export const callDeleteUser = (id: string) => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationKey: ['deleteUser', id],
    mutationFn: () => deleteUserById(id),
    onSuccess: () => {
      enqueueSnackbar({ message: 'Użytkownik usunięty.' });
      queryClient.invalidateQueries({ queryKey: ['user', id] });
    },
    onError: () => {
      enqueueSnackbar({ variant: 'error', message: 'Błąd podczas usuwania użytkownika.' });
    }
  });
};
