describe("Message format", () => {
  it("should match expected structure", () => {
    const message = {
      from: "ts-service",
      to: "go-service",
      message: "Hello!",
      date: new Date().toISOString()
    };

    expect(message).toHaveProperty("from");
    expect(message).toHaveProperty("to");
    expect(message).toHaveProperty("message");
    expect(message).toHaveProperty("date");
  });
});
