import { ProviderLoginButton } from "components/ProviderLoginButton";
import { useNavigate, useSearchParams } from "react-router";
import { LOGIN_PROVIDERS } from "constants/loginProviders";
import { useTranslation } from "react-i18next";
import { Button } from "lib/components/Button";
import { Card } from "lib/components/Card";
import { useEffect } from "react";

export function LoginPage() {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const { t } = useTranslation();

  useEffect(() => {
    if (document.cookie.includes("Authorization=Bearer")) {
      navigate("/");
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleClickReturn = () => {
    navigate("/");
  };

  return (
    <div className="h-full flex flex-col items-center justify-center gap-4">
      {searchParams.get("banned") ? (
        <Card>
          <h3 className="text-xl text-text-100">{t("youAreBannedFromApp")}</h3>
        </Card>
      ) : (
        <h1 className="text-2xl text-text-100">{t("applicationName")}</h1>
      )}

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
