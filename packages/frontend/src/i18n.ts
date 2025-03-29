import LanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";
import ptBR from "../i18n/ptBR.json";
import en from "../i18n/en.json";
import i18n from "i18next";

i18n
  .use(initReactI18next)
  .use(LanguageDetector)
  .init({
    fallbackLng: "en",

    detection: {
      convertDetectedLanguage: (lng) => lng.replace("-", ""),
    },

    resources: {
      en,
      ptBR,
    },
  });

export default i18n;
