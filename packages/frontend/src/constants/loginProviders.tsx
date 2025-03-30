import { ReactNode } from "react";
import {
  FaDiscord,
  FaFacebook,
  FaGithub,
  FaGoogle,
  FaInstagram,
} from "react-icons/fa";

interface LoginProvider {
  provider: string;
  icon: ReactNode;
  bgClassName: string;
}

export const LOGIN_PROVIDERS: LoginProvider[] = [
  {
    provider: "Discord",
    icon: <FaDiscord />,
    bgClassName: "bg-[#5865F240]",
  },
  {
    provider: "Facebook",
    icon: <FaFacebook />,
    bgClassName: "bg-[#1877F240]",
  },
  {
    provider: "Github",
    icon: <FaGithub />,
    bgClassName: "bg-[#E6E6E640]",
  },
  {
    provider: "Google",
    icon: <FaGoogle />,
    bgClassName: "bg-[#E3362940]",
  },
  {
    provider: "Instagram",
    icon: <FaInstagram />,
    bgClassName: "bg-[#EE4C5E40]",
  },
];
