'use client';

import React from 'react';
import { Formik, Field, Form, FormikHelpers, ErrorMessage } from 'formik';
import useAuth, { User } from './hooks/useAuth';


const Login = () => {
  const { mutateAsync: signInMutation, } = useAuth();
  return (
    <div className="flex items-center justify-center min-h-screen">
      <div className="bg-white p-8 rounded shadow-md w-96">
        <h1 className="text-2xl font-bold mb-4">Login</h1>
        <Formik
          initialValues={{
            password: '',
            email: '',
          }}
          onSubmit={async (values: User, { setSubmitting, setErrors }: FormikHelpers<User>) => {
            try {
              await signInMutation(values);
            } catch (error) {
              const { message } = error as Error;
              setErrors({ password: message });
            } finally {
              setSubmitting(false);
            }
          }}
        >
          <Form>
            <div className="mb-4">
              <label htmlFor="email" className="block text-sm font-medium text-gray-600">
                Email
              </label>
              <Field type="text" id="email" name="email" className="input input-bordered w-full max-w-xs" />
            </div>
            <div className="mb-4">
              <label htmlFor="password" className="block text-sm font-medium text-gray-600">
                Password
              </label>
              <Field type="password" id="password" name="password" className="input input-bordered w-full max-w-xs" />
              <ErrorMessage name="password" component="div" className="alert alert-error" />
            </div>
            <button type="submit" className="btn btn-primary">
              Submit
            </button>
          </Form>
        </Formik>
      </div>
    </div>
  );
};

export default Login;
