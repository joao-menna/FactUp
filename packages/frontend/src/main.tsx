import { RouterProvider } from "react-router";
import { createRoot } from "react-dom/client";
import { StrictMode } from "react";
import { router } from "router";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
