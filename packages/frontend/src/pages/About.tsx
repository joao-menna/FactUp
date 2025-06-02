import { useTranslation } from "react-i18next";
import { FaGithub } from "react-icons/fa";
import { clsx } from "clsx/lite";

export function AboutPage() {
  const { t } = useTranslation();
  return (
    <div
      className={clsx(
        "size-full flex flex-col justify-center items-center",
        "text-text-200 gap-4"
      )}
    >
      <h1 className={clsx("text-4xl")}>{t("applicationName")}</h1>
      <p className={clsx("text-lg max-w-96")}>{t("aboutPageText")}</p>
      <a
        target="_blank"
        referrerPolicy="no-referrer"
        href="https://github.com/joao-menna/FactUp"
      >
        <FaGithub className="size-12 hover:text-text-200/80 duration-150" />
      </a>
    </div>
  );
}
