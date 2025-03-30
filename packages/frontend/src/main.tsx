import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { LoginProvider } from "providers/LoginProvider";
import { RouterProvider } from "react-router";
import { createRoot } from "react-dom/client";
import { StrictMode } from "react";
import { router } from "router";
import "./styles.css";
import "./i18n";

const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <LoginProvider>
        <RouterProvider router={router} />
      </LoginProvider>
    </QueryClientProvider>
  </StrictMode>
);
