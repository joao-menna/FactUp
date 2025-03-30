import { Sidebar } from "components/Sidebar";
import { Outlet } from "react-router";

export function MainLayout() {
  return (
    <div className="h-full bg-primary-900 md:flex">
      <Sidebar />
      <Outlet />
    </div>
  );
}
