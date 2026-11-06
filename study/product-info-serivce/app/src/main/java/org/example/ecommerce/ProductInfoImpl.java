package org.example.ecommerce;

import ecommerce.ProductInfoGrpc;
import ecommerce.ProductInfoOuterClass;
import ecommerce.ProductInfoOuterClass.*;
import ecommerce.ProductInfoOuterClass.Product;
import io.grpc.Status;
import io.grpc.StatusException;
import io.grpc.stub.StreamObserver;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class ProductInfoImpl extends ProductInfoGrpc.ProductInfoImplBase{

    private Map productMap = new HashMap<String, Product>();

    @Override
    public void getProduct(ProductID request, StreamObserver<Product> responseObserver) {
        String id = request.getValue();
        if(productMap.containsKey(id)) {
            responseObserver.onNext((Product) productMap.get(id));
            responseObserver.onCompleted();;
        } else responseObserver.onError(new StatusException(Status.NOT_FOUND));
    }

    @Override
    public void addProduct(Product request, StreamObserver<ProductID> responseObserver) {
        UUID uuid = UUID.randomUUID();
        String randomUUIDString = uuid.toString();
        request = request.toBuilder().setId(randomUUIDString).build();
        productMap.put(randomUUIDString, request);
        ProductInfoOuterClass.ProductID id = ProductInfoOuterClass.ProductID.newBuilder().setValue(randomUUIDString).build();
        responseObserver.onNext(id);
        responseObserver.onCompleted();
    }
}
