import { ProviderLoginButton } from "components/ProviderLoginButton";
import { LOGIN_PROVIDERS } from "constants/loginProviders";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router";
import { Button } from "lib/components/Button";
import { Card } from "lib/components/Card";

export function LoginPage() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const handleClickReturn = () => {
    navigate("/");
  };

  return (
    <div className="h-full flex flex-col items-center justify-center gap-4">
      <h1 className="text-2xl text-text-100">{t("applicationName")}</h1>

      <p className="text-white">{t("logInWith")}</p>

      <Card className="flex flex-col gap-2 w-52">
        {LOGIN_PROVIDERS.map((p) => (
          <ProviderLoginButton key={p.provider} {...p} />
        ))}
      </Card>

      <Button
        className="bg-accent-400 hover:bg-accent-400/80 text-lg"
        onClick={handleClickReturn}
      >
        {t("return")}
      </Button>
    </div>
  );
}
