import { z } from 'zod';

const userLoginSchema = z.object({
  email: z.string().email('Nieprawidłowy format adresu email'),
  password: z.string().min(6, 'Hasło musi mieć co najmniej 6 znaków')
});

export type UserLogin = z.infer<typeof userLoginSchema>;
export { userLoginSchema };
