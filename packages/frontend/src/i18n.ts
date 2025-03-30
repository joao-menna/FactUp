import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";
import ptBR from "../i18n/ptBR.json";
import enUS from "../i18n/enUS.json";
import i18n from "i18next";

i18n
  .use(initReactI18next)
  .use(LanguageDetector)
  .init({
    fallbackLng: "enUS",

    detection: {
      convertDetectedLanguage: (lng) => lng.replace("-", ""),
    },

    resources: {
      enUS: {
        translation: {
          ...enUS,
        },
      },
      ptBR: {
        translation: {
          ...ptBR,
        },
      },
    },
  });

export default i18n;
