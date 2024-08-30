import getConfigPath, { Operation } from "../config";

test("Running my first test", function () {
  const config = getConfigPath({});
  expect(config.operation).toEqual(Operation.Print);
  expect(config.args).toEqual([]);
});
test("print key", function () {
  const config = getConfigPath({
    args: ["foo"],
  });
  expect(config.operation).toEqual(Operation.Print);
  expect(config.args).toEqual(["foo"]);
});
test("should create an add projector config", function () {
  const config = getConfigPath({
    args: ["add", "foo", "bar"],
  });
  expect(config.operation).toEqual(Operation.Add);
  expect(config.args).toEqual(["foo", "bar"]);
});

test("should remove a projector from config", function () {
  const config = getConfigPath({
    args: ["rm", "foo"],
  });
  expect(config.operation).toEqual(Operation.Remove);
  expect(config.args).toEqual(["foo"]);
});
