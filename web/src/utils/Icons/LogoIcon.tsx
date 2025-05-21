const LogoIcon = (props: any) => (
  <svg xmlns="http://www.w3.org/2000/svg" width={300} height={300} {...props}>
    <defs>
      <radialGradient id="a" cx="50%" cy="30%" r="60%">
        <stop offset="0%" stopColor="#fc0" />
        <stop offset="100%" stopColor="transparent" />
      </radialGradient>
    </defs>
    <circle cx={150} cy={90} r={80} fill="url(#a)" />
    <path stroke="#fff" strokeWidth={8} d="m60 60 40 50M240 60l-40 50" />
    <circle cx={150} cy={130} r={30} fill="#fff" />
    <rect width={60} height={80} x={120} y={160} fill="#fff" rx={10} />
  </svg>
);
export default LogoIcon;
