import { HomePage } from "pages/Home";
import { RouteObject } from "react-router";

export const mainRoutes: RouteObject[] = [
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "u",
    children: [
      {
        path: ":userId",
      },
    ],
  },
  {
    path: "curiosities",
  },
  {
    path: "sayings",
  },
];
