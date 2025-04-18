import { createBrowserRouter } from "react-router";
import { LoginLayout, MainLayout } from "layouts";
import { ErrorPage } from "pages/Error";
import { Page404 } from "pages/Page404";
import { loginRoutes } from "./login";
import { mainRoutes } from "./main";

export const router = createBrowserRouter([
  {
    errorElement: <ErrorPage />,
    children: [
      {
        element: <MainLayout />,
        children: mainRoutes,
      },
      {
        element: <LoginLayout />,
        children: loginRoutes,
      },
      {
        path: "*",
        element: <Page404 />,
      },
    ],
  },
]);
