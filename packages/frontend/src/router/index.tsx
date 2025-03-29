import { createBrowserRouter } from "react-router";
import { LoginLayout, MainLayout } from "layouts";
import { loginRoutes } from "./login";
import { mainRoutes } from "./main";

export const router = createBrowserRouter([
  {
    element: <MainLayout />,
    children: mainRoutes,
  },
  {
    element: <LoginLayout />,
    children: loginRoutes,
  },
]);
