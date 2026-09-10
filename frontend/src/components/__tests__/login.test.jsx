import { render, screen } from "@testing-library/react"
import userEvent from "@testing-library/user-event"
import { MemoryRouter } from "react-router-dom"
import { AuthProvider } from "../../context/AuthContext"
import Login from "../login"
import { api } from "../../services/api"
import { vi } from "vitest"

vi.mock("../../services/api", () => ({
  api: {
    post: vi.fn(),
  },
}))

describe("Login", () => {
  beforeEach(() => {
    localStorage.clear()
    vi.clearAllMocks()
  })

  test("renders login form", () => {
    render(
      <MemoryRouter>
        <AuthProvider>
          <Login />
        </AuthProvider>
      </MemoryRouter>
    )

    expect(
      screen.getByLabelText("Email")
    ).toBeInTheDocument()

    expect(
      screen.getByLabelText("Password")
    ).toBeInTheDocument()

    expect(
      screen.getByRole("button", { name: /login/i })
    ).toBeInTheDocument()
  })

  test("allows user to enter email and password", async () => {
    const user = userEvent.setup()

    render(
      <MemoryRouter>
        <AuthProvider>
          <Login />
        </AuthProvider>
      </MemoryRouter>
    )

    const emailInput = screen.getByLabelText("Email")
    const passwordInput = screen.getByLabelText("Password")

    await user.type(
      emailInput,
      "test@example.com"
    )

    await user.type(
      passwordInput,
      "password123"
    )

    expect(emailInput).toHaveValue(
      "test@example.com"
    )

    expect(passwordInput).toHaveValue(
      "password123"
    )
  })

  test("logs in successfully", async () => {
    const user = userEvent.setup()

    api.post.mockResolvedValue({
      token: "test-jwt-token",
    })

    render(
      <MemoryRouter>
        <AuthProvider>
          <Login />
        </AuthProvider>
      </MemoryRouter>
    )

    await user.type(
      screen.getByLabelText("Email"),
      "test@example.com"
    )

    await user.type(
      screen.getByLabelText("Password"),
      "password123"
    )

    await user.click(
      screen.getByRole("button", { name: /login/i })
    )

    expect(api.post).toHaveBeenCalledWith(
      "/auth/login",
      {
        email: "test@example.com",
        password: "password123",
      }
    )

    expect(
      localStorage.getItem("token")
    ).toBe("test-jwt-token")
  })
})