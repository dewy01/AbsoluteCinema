import { Box, TextField, Button, Link, Paper, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { userLoginSchema, type UserLogin } from '@/schemas/UserLogin';
import { callUserLogin } from '@/apis/User';

export const LoginView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset
  } = useForm<UserLogin>({
    resolver: zodResolver(userLoginSchema),
    mode: 'onChange',
    defaultValues: {
      email: '',
      password: ''
    }
  });

  const { mutateAsync } = callUserLogin();

  const onSubmit = (data: UserLogin) => {
    mutateAsync(data, {
      onSuccess: () => {
        reset();
      },
      onError: (error) => {
        console.error('Login failed:', error);
        reset();
      }
    });
  };

  return (
    <Box
      component="section"
      sx={{
        display: 'flex',
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%'
      }}>
      <Paper
        elevation={6}
        sx={{
          p: 4,
          borderRadius: 3,
          backgroundColor: 'background.paper',
          minWidth: '300px'
        }}>
        <Typography variant="h5" component="h1" align="center" sx={{ mb: 2, fontWeight: 'bold' }}>
          Zaloguj się
        </Typography>

        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            '& .MuiTextField-root': { m: 1, width: '25ch' }
          }}
          noValidate
          autoComplete="off">
          <TextField
            label="Adres email"
            {...register('email')}
            error={!!errors.email}
            helperText={errors.email?.message}
          />

          <TextField
            label="Hasło"
            type="password"
            {...register('password')}
            error={!!errors.password}
            helperText={errors.password?.message}
          />

          <Button
            type="submit"
            variant="contained"
            color="primary"
            sx={{ mt: 2, width: '25ch' }}
            disabled={isSubmitting}>
            Zaloguj się
          </Button>

          <Link href="#" underline="hover" sx={{ mt: 1, fontSize: '0.875rem' }}>
            Nie pamiętasz hasła?
          </Link>
        </Box>
      </Paper>
    </Box>
  );
};
