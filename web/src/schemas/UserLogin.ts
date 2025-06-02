import { z } from 'zod';

const userLoginSchema = z.object({
  email: z.string().min(1, 'Email jest wymagany').email('Nieprawidłowy format adresu email'),
  password: z.string().min(1, 'Hasło jest wymagane').min(6, 'Hasło musi mieć co najmniej 6 znaków')
});

export type UserLogin = z.infer<typeof userLoginSchema>;
export { userLoginSchema };
