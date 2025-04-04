import { Button } from "lib/components/Button";
import { ReactNode } from "react";
import { clsx } from "clsx/lite";

interface Props {
  provider: string;
  bgClassName?: string;
  icon: ReactNode;
}

export function ProviderLoginButton({ icon, provider, bgClassName }: Props) {
  const handleClickProviderLogin = () => {
    const baseUrl = import.meta.env.VITE_BACKEND_BASE_URL ?? "";
    const providerName = provider.toLowerCase();

    location.href = `${baseUrl}/api/v1/auth/login/${providerName}`;
  };

  return (
    <Button
      className={clsx(
        "flex gap-2 items-center px-2 py-1.5",
        "text-lg justify-center",
        bgClassName
      )}
      onClick={handleClickProviderLogin}
    >
      <div>{icon}</div>
      <span>{provider}</span>
    </Button>
  );
}
