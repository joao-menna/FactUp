interface User {
  id: number;
  imagePath: string;
  createdAt: string;
  displayName: string;
}

class UserService {
  baseUrl = `${import.meta.env.VITE_BACKEND_BASE_URL ?? ""}/api/v1/user`;

  async getLogged() {
    const req = await fetch(this.baseUrl, {
      credentials: "include",
    });

    const json = (await req.json()) as User;

    return json;
  }

  async getById(userId: number) {
    const req = await fetch(`${this.baseUrl}/${userId}`, {
      credentials: "include",
    });

    const json = (await req.json()) as User;

    return json;
  }
}

export const userService = new UserService();
