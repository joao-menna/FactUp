import { Outlet } from "react-router";

export function LoginLayout() {
  return (
    <div className="h-full bg-primary-900">
      <Outlet />
    </div>
  );
}
