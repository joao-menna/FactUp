import { LoginCallback } from "pages/LoginCallback";
import { RouteObject } from "react-router";
import { LoginPage } from "pages/Login";

export const loginRoutes: RouteObject[] = [
  {
    path: "login",
    element: <LoginPage />,
  },
  {
    path: "login/callback",
    element: <LoginCallback />,
  },
];
