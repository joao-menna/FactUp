import { useTranslation } from "react-i18next";
import { clsx } from "clsx/lite";

interface Props {
  className?: string;
  imagePath: string;
}

export function ProfilePicture({ className = "size-12", imagePath }: Props) {
  const { t } = useTranslation();

  return (
    <div className={clsx(className)}>
      <img
        className={clsx("rounded-lg")}
        src={imagePath}
        alt={t("profilePicture")}
      />
    </div>
  );
}
