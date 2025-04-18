import { PostListPage } from "pages/PostList";
import { SendPostPage } from "pages/SendPost";
import { RouteObject } from "react-router";
import { HomePage } from "pages/Home";
import { PostPage } from "pages/Post";

export const mainRoutes: RouteObject[] = [
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "p",
    children: [
      {
        path: ":postId",
        element: <PostPage />,
      },
    ],
  },
  {
    path: "u",
    children: [
      {
        path: ":userId",
        element: <PostListPage type="both" forPage="user" />,
      },
    ],
  },
  {
    path: "post",
    element: <SendPostPage />,
  },
  {
    path: "facts",
    element: <PostListPage type="fact" />,
  },
  {
    path: "sayings",
    element: <PostListPage type="saying" />,
  },
];
