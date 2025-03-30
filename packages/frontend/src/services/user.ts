class UserService {
  baseUrl = `${import.meta.env.VITE_BACKEND_BASE_URL}/api/v1/user`;

  async getLogged() {
    const req = await fetch(this.baseUrl);

    const json = await req.json();

    return json;
  }

  async getById(userId: number) {
    const req = await fetch(`${this.baseUrl}/${userId}`);

    const json = await req.json();

    return json;
  }
}

export const userService = new UserService();
