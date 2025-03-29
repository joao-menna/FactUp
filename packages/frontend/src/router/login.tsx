import { RouteObject } from "react-router";

export const loginRoutes: RouteObject[] = [
  {
    path: "login",
    children: [
      {
        path: "callback",
      },
    ],
  },
];
